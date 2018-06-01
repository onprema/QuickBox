function displaySwitch (data) {

  // Hide loading section
  $('.loading').hide()

  // Display the running container after clicking 'start'
  $('.running-container').show()

  // Display container ID and Destroy button after clicking 'start'
  $("ul#running-container").append(
    '<li>' + 'id: ' + String(data.Id) + '</li>' +
    '<li>' + 'port: ' + String(data.Port) + '</li>' +
    '<li>' + 'file: ' + String(data.Dockerfile) + '</li>' +
    '<li>' + 'name: ' + String(data.Name) + '</li>' +
    '<button id="destroy">Destroy</button>')
}
