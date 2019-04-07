import $ from 'jquery'
import axios from 'axios'
import * as THREE from 'three'
import OrbitControls from 'three-orbitcontrols'


const onImage = new THREE.MeshBasicMaterial({ opacity: .8, transparent: true, color: 'rgba(0,200,0)' });
const offImage = new THREE.MeshBasicMaterial({ opacity: 0.1, transparent: true, color: 'rgba(0,190,0)' });
const offMiniColumn = new THREE.MeshBasicMaterial({ opacity: 0.5, transparent: true, color: 'rgba(40,100,40)' });


function learnLive(image) {
  axios.get('http://localhost:3000/learnings/' + image)
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

function draw(object) {
  let spatialPooler = object.spatialPooler
  let encoded = object.encoded
  let image = object.image
  let w = object.inputSpaceWidth
  let h = object.inputSpaceHeight
  let spSquare = parseInt(Math.sqrt(spatialPooler.neurons.length))
  let threshold = object.threshold
  let overlap = object.overlap


  const geometryImage = new THREE.BoxGeometry(0.09, 0.09, 0.09);
  const geometryNeuron = new THREE.BoxGeometry(0.09, 0.09, 0.09);
  const geometryMiniColumNeuron = new THREE.BoxGeometry(0.09, 0.09, 0.09);

  var scene = new THREE.Scene();
  var camera = new THREE.PerspectiveCamera(75, window.innerWidth / window.innerHeight, 0.1, 1000);
  camera.position.z = 10;
  camera.position.x = 0;
  // camera.position.x = Math.PI * -1;
  var renderer = new THREE.WebGLRenderer();
  renderer.setSize(window.innerWidth, window.innerHeight);
  document.body.appendChild(renderer.domElement);

  // draw cup
  for (var i = 0; i < image.length; i++) {
    const x = (i % w)
    const y = parseInt(i / w)
    var cube = new THREE.Mesh(geometryImage, image.charAt(i) === 'X' ? onImage : offImage);
    cube.position.x = x * 0.1 - (w * 0.5 * 0.1)
    cube.position.y = y * -0.1 + (h * 0.5 * 0.1)
    cube.position.z = 1
    scene.add(cube);
  }

  // draw encoded
  for (var i = 0; i < encoded.length; i++) {
    const x = (i % w)
    const y = parseInt(i / w)
    var cube = new THREE.Mesh(geometryImage, encoded.charAt(i) === 'X' ? onImage : offImage);
    cube.position.x = x * 0.1 - (w * 0.5 * 0.1)
    cube.position.y = y * -0.1 + (h * 0.5 * 0.1)
    cube.position.z = 0
    scene.add(cube);
  }

  // draw Spatial Pooler Neurons
  let spWidth = (w * 0.1) + 0.2
  let spHeight = (h * 0.1) + 0.2
  let initialX = -(spWidth * spSquare * 0.5) + (w * 0.5 * 0.1)
  let initialY = (spHeight * spSquare * 0.5) + (h * 0.5 * 0.1)
  let currentXOffset = initialX
  let currentYOffset = initialY
  let count = 0
  spatialPooler.neurons.forEach(neuron => {
    let minicount = 0
    neuron.miniColumnNeurons.forEach(mn => {
      var cube = new THREE.Mesh(geometryMiniColumNeuron, offMiniColumn);
      cube.position.x = currentXOffset - (w * 0.5 * 0.1) + minicount*0.1
      cube.position.y = currentYOffset - (h * 0.5 * 0.1) + 0.1
      cube.position.z = -1
      scene.add(cube);
      minicount++
    })
    neuron.proximalInputs.forEach(d => {
      let permanence = d.permanence
      let coord = d.connectedNeuronId
      let nx = parseInt(coord.substring(1).split(',')[0])
      let ny = parseInt(coord.substring(1, coord.length - 1).split(',')[1])
      let permanenceColour = parseInt(155 + (100 / 9) * permanence)
      let activeGreen = 50
      if (neuron.active) {
        activeGreen = permanenceColour
        permanenceColour = 50
      }
      let material = new THREE.MeshBasicMaterial({ opacity: 0.3, transparent: true, color: 'rgba(50,' + activeGreen + ',' + permanenceColour + ')' });
      if (encoded.charAt(ny * w + nx) === 'X') {
        if (permanence >= threshold) {
          material = new THREE.MeshBasicMaterial({ opacity: 0.8, transparent: true, color: 'rgba(200,0,0)' });
        } else {
          material = new THREE.MeshBasicMaterial({ opacity: 0.8, transparent: true, color: 'rgba(200,200,200)' });
        }
      }
      var cube = new THREE.Mesh(geometryNeuron, material);
      cube.position.x = currentXOffset + (nx * 0.1 - (w * 0.5 * 0.1))
      cube.position.y = currentYOffset + (ny * -0.1 - (h * 0.5 * 0.1))
      cube.position.z = -1
      scene.add(cube);

    })


    count++
    currentXOffset += spWidth
    if (count % spSquare == 0) {
      currentYOffset -= spHeight
      currentXOffset = initialX
    }
  })

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