// Verify Public Key (TODO)
function verifyKey(key) {
  if (key.length < 20) {
    return null
  } else {
    creds['key'] = key
  } return 1
}

function verifyGithubUsername() {
  $('.message p').show()
  let username = $('input#username').val()
  let url = 'https://api.github.com/users/' + username
  // Validate github username
  $('div#loader').show()
  $.ajax ({
    type: 'GET',
    url: url,
    statusCode: {
      404: function() {
        $('div#loader').hide()
        $('.message').text('Invalid GitHub username. Please try again')
        $('.register').show()
        $('.register p').text('')
      }
    },
    success: function (data) {
      // Add username to creds object
      creds['username'] = username;
      $('form#registration').hide()
      $('span#username').html(username)
      $('div#userSuccess').delay(200).fadeIn('normal', function() {
        $(this).show()
      })
      $('.message p').text('Welcome, ' + username + '!')
    }
  })
 
  // Check if user has public keys linked to the github account
  $.ajax ({
    type: 'GET',
    url: url + '/keys',
    success: function (data) {
      $('div#loader').hide()
      if (data.length > 0) {
        let keys = []
        for (let i = 0; i < data.length; i++) {
          keys.push(data[i]['key'])
        }

        // Add keys to creds object
        creds['keys'] = keys;

        // Display the github key and confirm they want to use them
        $('div#keyMsgSuccess').delay(500).fadeIn('normal', function() {
          $(this).show()
        })

        // If they click yes, send key to backend and display service
        $('button#keyConfirmYes').click(function (event) {
          event.preventDefault()

          // send public key to backend
          let credURL = '/api/v1/register'
          $.ajax ({
            type: 'GET',
            url: credURL,
            headers: {
              creds: JSON.stringify(creds['keys'])
            },
            success: function(data) {
              console.log(data)
            }
          })

          // Display the service
          $('div#loader').hide()
          $('div#keyMsgSuccess').hide()
          $('.service').show()

          // Set cookie?
					Cookies.set(username, true, { expires: .01 })
          console.log(document.cookie)
          $('span#verified').text('[verified]')
        })

        // Alternatively, prove an SSH key
        $('button#customKey').click( function(event) {
          event.preventDefault()
          let key = $('textarea#customKey').val()
          let result = verifyKey(key);
          if (result != null) {
            $('div#loader').hide()
            $('div#keySuccess').delay(500).fadeIn('normal', function() {
              $(this).text('Key added successfully.')
            })
            creds['keys'] = key
            $('div#keyMsg').hide()

            // send public key to backend
            let credURL = '/api/v1/register'
            $.ajax ({
              type: 'GET',
              url: credURL,
              headers: {
                creds: JSON.stringify(creds['keys'])
              },
              success: function(data) {
                console.log(data)
              }
            })
            $('div#keySuccess').hide()
            $('.service').show()
          } else {
            $('div#badCustomKeyFlash').delay(500).fadeIn('normal', function() {
              $(this).text('Invalid key. Please verify and try again.')
              setTimeout(function () {
                location.reload()
              }, 5000)
            })
          }
        })
      } else {
        $('div#keyMsg').delay(500).fadeIn('normal', function() {
          $('div#loader').hide()
          $(this).show()
          //setTimeout(location.reload(), 5000)
        })
        }
      }
   })
  }
