// Global variables
let containerSpecs = {};
let cloneURL = '';
let creds = {};

$(document).ready(function () {
  // Collect the base image for the container to be ran (required)
  $('input#base').on('click', function () {
    if ($(this).prop('checked')) {
      containerSpecs[$(this).attr('data-key')] = $(this).attr('data-val');
    } else {
      delete containerSpecs[$(this).attr('data-key')];
    }
  });

  // Make the default base ubuntu:14.04
  if (containerSpecs['base'] === undefined) {
    containerSpecs['base'] = 'qb-ubuntu:14.04';
  }

  // Add random root password
  containerSpecs['pw'] = hash();

  // Make a GET request to send container specs to API when 'Start' is clicked
  $('input#start').click(function (event) {
    event.preventDefault();
    $('.service').hide();
    $('.footer').hide();
    $('#loader').show();

    // Validate github repo
    let val = $('input#repo-input').val();
    if (val !== '') {

      // Remove trailing slash, or `.git` if the enter clone URL
      if (val[val.length-1] == '/') { val = val.slice(0, -1) };
      if (val.slice(-4) == '.git') { val = val.slice(0, -4) };

      $('.message p').text('Validating repository...')

      let len = 'https://github.com/'.length;
      let repo = val.slice(len);
      let url = 'https://api.github.com/repos/' + repo;

      $.ajax({
        type: 'GET',
        url: url,
        statusCode: {
          // If repo is not valid GitHub API returns 404
          404: function () {
            $('#loader').hide();
            $('.service').show();
            $('.footer').show();
            $('.message p').html('<p style="color: red">Invalid GitHub repository.</p>')
          }
        },
        // If repo is valid GitHub API returns 200
        success: function (data) {
          $('.message p').fadeIn('normal', function () {
            $(this).html('<p style="color: green">GitHub repository validated.</p>')
          });

          // Add repo to containerSpecs (these are global variables)
          cloneURL = data.clone_url;
          // Replacing '/' with '|' to pass as url query string
          containerSpecs['cloneURL'] = cloneURL.replace(/\//g, '|');
          containerSpecs['name'] = data.name.toLowerCase();
          containerSpecs['id'] = String(data.id);

          // Start build...
          // startContainer is found in `functions.js`
          startContainer();
        }
      });
    } else {
      // If they don't enter a repository just build the base image
      startContainer();
    }
  });
});
