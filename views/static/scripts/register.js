let creds = {}
$(document).ready(function () {
  // Hide service until users are verified
  $('.service').hide()
  $('div#keyMsgSuccess').hide()

  // Username submit click
  $('button#username').click(function (event) {
    event.preventDefault();
    console.log('here')
    verifyGithubUsername();
  })
})
