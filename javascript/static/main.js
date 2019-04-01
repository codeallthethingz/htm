import objects from './objects/*.yml'  // loads all the yaml objects and parses them.
import $ from "jquery"
import axios from "axios"

$(() => {

  const cellWidth = 8
  const cellHeight = 16
    const $container = $('#environments')

    function renderObjectToCanvas(object, $parent) {
        let $title = $('<div>').html(object.name)
        let $canvas = $('<canvas>')

        $parent.append($title)
        $parent.append($canvas)

        const ctx = $canvas[0].getContext('2d')
        $canvas.attr('width', (cellWidth * object.InputSpaceWidth + 30 ) * object.Cells.length)
        $canvas.attr('height', cellHeight * object.InputSpaceHeight)
        ctx.font = cellHeight + 'px sans-serif'
        var offset = cellWidth * object.InputSpaceWidth + 30
        var currentOffset = 0
        object.Cells.forEach(cell => {
          cell.Coordinates.forEach(coord => {
            var permenance = cell.Permenances[cell.CoordLookup[coord]]
            ctx.fillStyle = 'rgba(' + (10*permenance + 150) + ',50,50,' + permenance / 10 + ')'
            ctx.fillRect((coord % object.InputSpaceWidth) * cellWidth + currentOffset, parseInt(coord / object.InputSpaceHeight ) * cellHeight - cellHeight, cellWidth, cellHeight)
            ctx.fillStyle = '#000000'
            ctx.fillText(
              permenance > 6 ? "x" : " ",
              (coord % object.InputSpaceWidth) * cellWidth + currentOffset,
              parseInt(coord / object.InputSpaceHeight ) * cellHeight -2)
            })
            currentOffset += offset
        })
    }

    // render all the objects
    Object.values(objects).forEach((value) => {
        renderObjectToCanvas(value, $container)
    })

    
    axios.get('http://localhost:3000/spatialPooler')
    .then(function (response) {
      // handle success
      console.log(response.data)
      renderObjectToCanvas(response.data, $container)
    })
  
    // Or render them individually
    // renderObjectToCanvas(objects['a'], $container)

})
