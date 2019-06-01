import $ from 'jquery'
import axios from 'axios'
import * as THREE from 'three'
import OrbitControls from 'three-orbitcontrols'


const drawSpatialPoolerFineGrained = false
const drawImage = false

const onImage = new THREE.MeshBasicMaterial({ opacity: .8, transparent: true, color: 'rgba(0,200,0)' });
const offImage = new THREE.MeshBasicMaterial({ opacity: 0.4, transparent: true, color: 'rgba(0,190,0)' });
const offMiniColumn = new THREE.MeshBasicMaterial({ opacity: 0.2, transparent: true, color: 'rgba(100,100,100)' });
const activeMiniColumn = new THREE.MeshBasicMaterial({ opacity: 0.9, transparent: true, color: 'rgba(40,255,40)' });
const predictiveMiniColumn = new THREE.MeshBasicMaterial({ opacity: 0.5, transparent: true, color: 'rgba(0,0,255)' });
const previouslyPredictiveMiniColumn = new THREE.MeshBasicMaterial({ opacity: 0.5, transparent: true, color: 'rgba(255,0,0)' });

function learnLive(image) {
  axios.get('http://localhost:3000/learnings/' + image + '?rand=' + Math.random())
    .then(function (response) {
      draw(response.data)
    })
}

function learn() {
  axios.get('test.data')
    .then(function (response) {
      draw(response.data)
    })
}

learn()
// learnLive('cup')

$('#console').on('keydown', e => {
  if (e.originalEvent.key == 'Enter') {
    var index = parseInt($('#console').val())
    console.log(index)
    let neuron = spatialPooler.neurons[index]
    console.log(neuron)
  }
})

let spatialPooler = {}
function draw(object) {


  spatialPooler = object.spatialPooler
  let encoded = object.encoded
  let image = object.image
  let w = object.inputSpaceWidth
  let h = object.inputSpaceHeight
  let spSquare = parseInt(spatialPooler.neurons.length / h)
  let threshold = object.threshold

  const geometryImage = new THREE.BoxGeometry(0.09, 0.09, 0.09);
  const g = 0.01
  const geometryNeuron = new THREE.BoxGeometry(g, g, g);
  const geometryMiniColumNeuron = new THREE.BoxGeometry(g, g, g);
  let active = 0
  let predictive = 0
  let previouslyPredictive = 0
  var scene = new THREE.Scene();
  var camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
  camera.position.z = 3
  // camera.z = 0
  camera.x = 0
  // camera.y = 0
  var renderer = new THREE.WebGLRenderer();
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);



  renderer.domElement.addEventListener("click", onclick, true);
  var selectedObject;
  var raycaster = new THREE.Raycaster();
  let allObjects = []
  function onclick(event) {
    var mouse = new THREE.Vector2();
    raycaster.setFromCamera(mouse, camera);
    var intersects = raycaster.intersectObjects(allObjects, true); //array
    if (intersects.length > 0) {

      intersects[0].face.color.setRGB(0.8 * Math.random() + 0.2, 0, 0);

      intersects[0].object.geometry.colorsNeedUpdate = true;
      console.log(intersects[0])
      animate()
    }
  }


  var currentXOffset = 0
  var currentYOffset = 0

  // draw minicolumns
  currentXOffset = 0
  currentYOffset = 1
  var c = 0
  spatialPooler.neurons.forEach(neuron => {

    let miniN = [].concat(neuron, neuron.miniColumnNeurons)
    let miniCount = 0
    miniN.forEach(n => {
      const x = (c % spSquare)
      const y = parseInt(c / spSquare)

      var mat = (n.previouslyPredictive && n.active) ? previouslyPredictiveMiniColumn :
        n.active ? activeMiniColumn :
          n.predictive ? predictiveMiniColumn : offMiniColumn
      var cube = new THREE.Mesh(geometryImage, mat);
      cube.position.x = -(x * 0.2)
      cube.position.y = (y * -0.1) - currentYOffset
      cube.position.z = -miniCount * 0.1
      scene.add(cube);
      miniCount++
      if (n.active) {
        active++;
      }
      if (n.predictive) {
        predictive++
      }
      if (n.previouslyPredictive) {
        previouslyPredictive++
      }
    })


    c++
  })

  $('#info').text('threshold: ' + threshold +
    ' active: ' + active +
    ' predictive: ' + predictive +
    ' previously predictive: ' + previouslyPredictive)

  // draw image
  if (drawImage) {
    currentXOffset = (w * 3 * 0.1) + (0.1 * 2)
    currentYOffset = (h * 0.5 * 0.1)
    for (var i = 0; i < image.length; i++) {
      const x = (i % w)
      const y = parseInt(i / w)
      var cube = new THREE.Mesh(geometryImage, image.charAt(i) === 'X' ? onImage : offImage);
      cube.position.x = x * 0.1 - currentXOffset
      cube.position.y = y * -0.1 + currentYOffset
      cube.position.z = 0
      scene.add(cube);
    }
  }

  // draw encoded
  currentXOffset = (w * 2 * 0.1) + (0.1 * 1)
  currentYOffset = (h * 0.5 * 0.1)
  for (var i = 0; i < encoded.length; i++) {
    const x = (i % w)
    const y = parseInt(i / w)
    var cube = new THREE.Mesh(geometryImage, encoded.charAt(i) === 'X' ? onImage : offImage);
    cube.position.x = x * 0.1 - currentXOffset
    cube.position.y = y * -0.1 + currentYOffset
    cube.position.z = 0
    allObjects.push(cube)
    scene.add(cube);
  }

  // draw Spatial Pooler Neurons
  if (drawSpatialPoolerFineGrained) {
    let spWidth = (w * g) + 0.02
    let spHeight = (h * g) + 0.02
    let initialX = -(w * 0.5 * 0.1)
    let initialY = (spHeight * spSquare * 0.5) + (h * 0.5 * g)
    currentXOffset = initialX
    currentYOffset = initialY
    let count = 0
    spatialPooler.neurons.forEach(neuron => {
      neuron.proximalInputs.forEach(d => {
        let permanence = d.p
        let coord = d.cId
        let nx = parseInt(coord.substring(1).split(',')[0])
        let ny = parseInt(coord.substring(1, coord.length - 1).split(',')[1])
        let permanenceColour = parseInt(100 + (50 / 9) * permanence)
        let activeGreen = 50
        if (neuron.active) {
          activeGreen = permanenceColour
          permanenceColour = 50
        }
        let material = new THREE.MeshBasicMaterial({ opacity: 0.8, transparent: true, color: 'rgba(50,' + activeGreen + ',' + permanenceColour + ')' });
        if (encoded.charAt(ny * w + nx) === 'X') {
          if (permanence >= threshold) {
            material = new THREE.MeshBasicMaterial({ opacity: 0.8, transparent: true, color: 'rgba(0,200,0)' });
          }
        }
        var cube = new THREE.Mesh(geometryNeuron, material);
        cube.position.x = (nx * g - (w * 0.5 * g)) - currentXOffset
        cube.position.y = (ny * -g - (h * 0.5 * g)) + currentYOffset
        cube.position.z = 0
        scene.add(cube);

      })


      count++
      currentXOffset += spWidth
      if (count % spSquare == 0) {
        currentYOffset -= spHeight
        currentXOffset = initialX
      }
    })
  }

  var controls = new OrbitControls(camera, renderer.domElement);
  controls.minDistance = 0;
  controls.maxDistance = 10;
  function animate() {
    requestAnimationFrame(animate);
    render();
  }
  function render() {
    renderer.render(scene, camera);
  }
  animate()
}