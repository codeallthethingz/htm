import $ from 'jquery'
import axios from 'axios'

$(() => {

  const cellWidth = 5
  const cellHeight = 10
  const poolerCellHeight = 10
  const $environmentActive = $('#environmentActive')

  function renderObjectToCanvas(object, $parent) {
    let spatialPooler = object.spatialPooler
    let encoded = object.encoded
    let image = object.image
    let threshold = object.threshold
    let overlap = object.overlap
    let $canvas = $('<canvas>')
    let spSquare = spatialPooler.neurons.length 
    $parent.append($canvas)
    const ctx = $canvas[0].getContext('2d')
    let canvasWidth = cellWidth * object.inputSpaceWidth
    let canvasHeight = cellHeight * object.inputSpaceHeight
    $canvas.attr('width', (canvasWidth + 30) * spSquare - 30)
    $canvas.attr('height', cellHeight * object.inputSpaceHeight + 10)
    ctx.font = cellHeight + 'px sans-serif'
    let xOffset = canvasWidth + 2 * cellWidth
    let yOffset = canvasHeight + 2 * cellHeight
    let currentXOffset = 0
    let currentYOffset = 0

    for (var i = 0; i < image.length; i++) {
      const x = (i % object.inputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / object.inputSpaceWidth) * cellHeight + currentYOffset;
      if (image.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    currentXOffset += xOffset
    for (var i = 0; i < encoded.length; i++) {
      const x = (i % object.inputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / object.inputSpaceWidth) * cellHeight + currentYOffset;
      if (encoded.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    currentXOffset += xOffset
    currentYOffset += 10
    let g = 0;
    spatialPooler.neurons.forEach(neuron => {
      ctx.fillText('Neuron ' + neuron.id + ', Score: ' + neuron.score, currentXOffset, currentYOffset)
      if (neuron.active) {
        ctx.fillStyle = 'rgba(0,0, 255, ' + Math.min(1, (neuron.score / 10) - 0.2) + ')'
        ctx.fillRect(currentXOffset, currentYOffset, canvasWidth, canvasHeight)
      }
      neuron.proximalInputs.forEach(dendrite => {
        let permanence = dendrite.p
        let coord = dendrite.cId
        let nx = parseInt(coord.substring(1).split(',')[0])
        let ny = parseInt(coord.substring(1, coord.length -1).split(',')[1])

        ctx.fillStyle = '#FFFFFF'
        if (permanence > threshold) {
          ctx.fillStyle = '#FFCCCC'
        }
        if (permanence > threshold && encoded && encoded.charAt(ny*object.inputSpaceWidth+nx) === 'X') {
          ctx.fillStyle = '#FF0000'
        }
        const x = nx * cellWidth + currentXOffset
        const y = ny * cellHeight + currentYOffset
        ctx.fillRect(x, y, cellWidth, cellHeight)
        ctx.fillStyle = '#000000'

        if (permanence > threshold && encoded && encoded.charAt(ny*object.inputSpaceWidth+nx) === 'X') {
          ctx.fillStyle = '#FFFFFF'
        }
        ctx.fillText(permanence == 0 ? "." : permanence > threshold ? permanence : permanence, x, y - 2 + cellHeight)
      })
      currentXOffset += xOffset
      g++
      if (g % spSquare == 0) {
        currentXOffset = 0
        currentYOffset += yOffset
      }
    })
    ctx.stroke()
  }
  function learn(image) {
    axios.get('http://localhost:3000/learnings/' + image)
      .then(function (response) {
        let $newDiv = $('<div>')
        $('#canvases').prepend($newDiv)
        renderObjectToCanvas(response.data, $newDiv)
      })
  }
  $('button').on('click', (e) => {learn(e.target.id);})
})