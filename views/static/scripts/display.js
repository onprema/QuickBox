function displaySwitch (data) {

  // Display the running container after clicking 'start'
  $('.running-container').show()

  // Hide the loader
  $('div#loader').hide()

  // Display Port and Destroy button after clicking 'start'
  $('.cmd').text('ssh root@myipaddress.com -p ' + String(data.Port))
  $('.running-container').append('<button id="destroy">Destroy</button>')
  $('.running-container').append('<p><span style="color: maroon; font-size: .85em">This container will automatically be destroyed in 12 hours.</p></span>')
}
