function displaySwitch (data) {

  // Display the running container after clicking 'start'
  $('.running-container').css('display', 'block')

  // Hide the original input fields
  $('.select-container').css('display', 'none')

  // Display container ID and Destroy button after clicking 'start'
  $("ul#running-container").append(
    '<li id="' + containerId + '">' + String(data.Id) + 
    '<button id="destroy">Destroy</button>' + '</li>')
}
