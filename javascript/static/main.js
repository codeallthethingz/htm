import $ from 'jquery'
import axios from 'axios'

$(() => {

  const cellWidth = 5
  const cellHeight = 10
  const poolerCellHeight = 10
  const $environmentActive = $('#environmentActive')

  function renderObjectToCanvas(object, $parent) {
    let spatialPooler = object.SpatialPooler
    let encoded = object.Encoded
    let image = object.Image
    let threshold = object.Threshold
    let overlap = object.Overlap
    let $canvas = $('<canvas>')
    let spSquare = spatialPooler.Cells.length 
    $parent.append($canvas)
    const ctx = $canvas[0].getContext('2d')
    let canvasWidth = cellWidth * object.InputSpaceWidth
    let canvasHeight = cellHeight * object.InputSpaceHeight
    $canvas.attr('width', (canvasWidth + 30) * spSquare - 30)
    $canvas.attr('height', cellHeight * spatialPooler.InputSpaceHeight + 10)
    ctx.font = cellHeight + 'px sans-serif'
    let xOffset = canvasWidth + 2 * cellWidth
    let yOffset = canvasHeight + 2 * cellHeight
    let currentXOffset = 0
    let currentYOffset = 0

    for (var i = 0; i < image.length; i++) {
      const x = (i % object.InputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / object.InputSpaceWidth) * cellHeight + currentYOffset;
      if (image.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    currentXOffset += xOffset
    for (var i = 0; i < encoded.length; i++) {
      const x = (i % object.InputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / object.InputSpaceWidth) * cellHeight + currentYOffset;
      if (encoded.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    currentXOffset += xOffset
    currentYOffset += 10
    let g = 0;
    spatialPooler.Cells.forEach(cell => {

      ctx.fillText('Cell ' + cell.ID + ', Score: ' + cell.Score, currentXOffset, currentYOffset)
      if (cell.Active) {
        ctx.fillStyle = 'rgba(0,0, 255, ' + Math.min(1, (cell.Score / 10) - 0.2) + ')'
        ctx.fillRect(currentXOffset, currentYOffset, canvasWidth, canvasHeight)
      }
      cell.Coordinates.forEach(coord => {
        let permanence = cell.Permanences[cell.CoordLookup[coord]]
        ctx.fillStyle = '#FFFFFF'
        if (permanence > threshold) {
          ctx.fillStyle = '#FFCCCC'
        }
        if (permanence > threshold && encoded && encoded.charAt(coord) === 'X') {
          ctx.fillStyle = '#7777FF'
        }
        const x = (coord % object.InputSpaceWidth) * cellWidth + currentXOffset
        const y = parseInt(coord / object.InputSpaceWidth) * cellHeight + currentYOffset
        ctx.fillRect(x, y, cellWidth, cellHeight)
        ctx.fillStyle = '#000000'
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

