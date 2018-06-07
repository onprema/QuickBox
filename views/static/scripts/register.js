$(document).ready(function () {
  // Hide service until users are verified
  $('.service').hide();

  // Username submit click
  $('button#username').click(function (event) {
    event.preventDefault();
    $('.message p').text('Verifying GitHub username...');
    let username = $('input#username').val();
    let invalidChars = ['!', '@', '$', '%', '^', '&', '*', '(', ')'];
    for (let i = 0; i < invalidChars.length; i++) {
      if (username.includes(invalidChars[i])) {
        $('div#loader').hide();
        $('.message p').html('<style="color:red">Invalid GitHub username. Please try again.</style>');
        $('.register').show();
      }
    }
    if (document.cookie.search(username) > -1 && username !== '') {
      $('.service').show();
      $('.message p').text('Welcome, ' + username + '!');
    } else {
      verifyGithubUsername();
    }
    $('.register').hide();
  });
});
