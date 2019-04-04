import $ from 'jquery'
import axios from 'axios'

$(() => {

  const cellWidth = 10
  const cellHeight = 10
  const pooler = 10

  function renderObjectToCanvas(object, $parent) {
    let spatialPooler = object.SpatialPooler
    let encoded = object.Encoded
    let image = object.Image
    let threshold = object.Threshold
    let imageWidth = spatialPooler.InputSpaceWidth
    let imageHeight = spatialPooler.InputSpaceHeight
    let overlap = object.Overlap
    let $canvas = $('<canvas>')
    let spSquare = parseInt(spatialPooler.Cells.length)
    $parent.append($canvas)
    const ctx = $canvas[0].getContext('2d')
    let canvasWidth = cellWidth * imageWidth
    let canvasHeight = cellHeight * imageHeight
    let xOffset = canvasWidth + 2 * pooler
    let yOffset = canvasHeight + 2 * pooler
    $canvas.attr('width', Math.min(40000, (canvasWidth * 2 + 30) + (spatialPooler.Cells.length) * pooler * imageWidth - 30))
    $canvas.attr('height', cellHeight * imageHeight + 10)
    ctx.font = cellHeight + 'px sans-serif'
    let currentXOffset = 0
    let currentYOffset = 0

    // raw image
    for (var i = 0; i < image.length; i++) {
      const x = (i % spatialPooler.InputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / spatialPooler.InputSpaceWidth) * cellHeight + currentYOffset;
      if (image.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }
    // encoded image
    currentXOffset += xOffset
    for (var i = 0; i < encoded.length; i++) {
      const x = (i % spatialPooler.InputSpaceWidth) * cellWidth + currentXOffset;
      const y = parseInt(i / spatialPooler.InputSpaceWidth) * cellHeight + currentYOffset;
      if (encoded.charAt(i) === 'X') {
        ctx.fillStyle = '#0000000'
        ctx.fillRect(x, y, cellWidth, cellHeight)
      }
    }

    // Spatial Pooler cells.
    currentXOffset += xOffset
    currentYOffset += 10
    let g = 0;
    ctx.font = '100px sans-serif'
    ctx.fillText('guess: ' + object.Guess, currentXOffset, currentYOffset + 200)
    ctx.font = cellHeight + 'px sans-serif'
    spatialPooler.Cells.forEach(cell => {

      ctx.fillStyle = '#000000'
      if (pooler > 3) {
        ctx.fillText('Cell ' + cell.ID + ', Score: ' + cell.Score, currentXOffset, currentYOffset)
      }
      if (cell.Active) {
        ctx.fillStyle = 'rgba(0,0, 255, ' + Math.min(1, (cell.Score / 10) - 0.2) + ')'
        ctx.fillRect(currentXOffset, currentYOffset, pooler * imageWidth, pooler * imageHeight)
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
        const x = (coord % spatialPooler.InputSpaceWidth) * pooler + currentXOffset
        const y = parseInt(coord / spatialPooler.InputSpaceWidth) * pooler + currentYOffset
        ctx.fillRect(x, y, pooler, pooler)
        if (pooler > 8) {
          ctx.fillStyle = '#000000'
          ctx.fillText(permanence == 0 ? "." : permanence > threshold ? permanence : permanence, x, y - 2 + cellHeight)
        }
      })
      currentXOffset += pooler * imageWidth
      g++
      if (g % spSquare == 0) {
        currentXOffset = xOffset + xOffset
        currentYOffset += pooler * imageHeight
      }
    })
    ctx.stroke()
  }
  function learn(image) {
    return axios.get('http://localhost:3000/learnings/' + image + '?rand=' + Math.random())
      .then(function (response) {
        let $newDiv = $('<div>')
        $('#canvases').prepend($newDiv)
        renderObjectToCanvas(response.data, $newDiv)
      })
  }
  $('.number').on('click', (e) => {
    let k = e.target.id.substring(1)
    learn5(k)
  })
  $('#reset').on('click', (e) => {
    axios.get('http://localhost:3000/learnings/reset')
  })
  let i = 1
  axios.get('http://localhost:3000/learnings/reset').then(() => {
    learn5(i)
  })
  function learn5(i) {
    learn(i).then(function () {
      learn(i).then(function () {
        learn(i).then(function () {
          learn(i).then(function () {
            learn(i)
          })
        })
      })
    })
  }
})