import $ from "jquery"
import axios from "axios"

$(() => {

  const cellWidth = 4
  const cellHeight = 6
  const poolerCellHeight = 10
  const $environmentActive = $('#environmentActive')

  function renderObjectToCanvas(object, $parent) {
    let spatialPooler = object.SpatialPooler
    let encoded = object.Encoded
    let image = object.Image
    let threshold = object.Threshold
    let overlap = object.Overlap
    let $canvas = $('<canvas>')
    let spSquare = parseInt(Math.sqrt(spatialPooler.Cells.length))
    $parent.append($canvas)
    const ctx = $canvas[0].getContext('2d')
    let canvasWidth = cellWidth * spatialPooler.InputSpaceWidth
    let canvasHeight = cellHeight * spatialPooler.InputSpaceHeight
    $canvas.attr('width', (canvasWidth + 30) * spSquare - 30)
    $canvas.attr('height', 6000)
    ctx.font = cellHeight + 'px sans-serif'
    let xOffset = canvasWidth + 2*cellWidth
    let yOffset = canvasHeight + 2*cellHeight
    let currentXOffset = 0
    let currentYOffset = 0

    for (var i = 0; i < image.length; i++) {
      const x = (i % spatialPooler.InputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / spatialPooler.InputSpaceWidth) * cellHeight + currentYOffset;
      if (image.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    currentXOffset += xOffset
    for (var i = 0; i < encoded.length; i++) {
      const x = (i % spatialPooler.InputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / spatialPooler.InputSpaceWidth) * cellHeight + currentYOffset;
      if (encoded.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    currentXOffset = 0
    currentYOffset = canvasHeight + 30
    let g = 0;
    spatialPooler.Cells.forEach(cell => {

      ctx.fillText("Cell " + cell.ID + ", Score: " + cell.Score + (cell.Active ? " (Active)" : ""), currentXOffset, currentYOffset)
      if (cell.Active) {
        ctx.fillStyle = "#7777FF"
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
        const x = (coord % spatialPooler.InputSpaceWidth) * cellWidth + currentXOffset
        const y = parseInt(coord / spatialPooler.InputSpaceWidth) * cellHeight + currentYOffset
        ctx.fillRect(x, y, cellWidth, cellHeight)
        ctx.fillStyle = '#000000'
        ctx.fillText(permanence > threshold ? permanence : permanence, x, y - 2 + cellHeight)
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

  axios.get('http://localhost:3000/activeForInput/cup')
    .then(function (response) {
      renderObjectToCanvas(response.data, $environmentActive)
    })

})
