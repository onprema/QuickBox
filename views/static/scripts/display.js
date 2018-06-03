function displaySwitch (data) {

  // Hide loading section
  $('#loader').hide()

  // Display the running container after clicking 'start'
  $('.running-container').show()

  // Display Port and Destroy button after clicking 'start'
  $("ul#running-container").append(
    '<li>' +'ssh root@mydomain.com -p ' + String(data.Port) + '</li>' +
    '<button id="destroy">Destroy</button>')
}
