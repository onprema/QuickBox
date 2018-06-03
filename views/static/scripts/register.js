let creds = {}

// Verify Public Key (TODO)
function verifyKey(key) {
  if (key.length < 20) {
    return null
  } else {
    creds['key'] = key
  } return 1
}

$(document).ready(function () {

  // Hide service until users are verified
  $('.service').hide()
  $('div#keyMsgSuccess').hide()

  // Username submit click
  $('button#username').on('click', function () {

    event.preventDefault();
    let username = $('input#username').val()
    let url = 'https://api.github.com/users/' + username

    // Validate github username
    $.ajax ({
      type: 'GET',
      url: url,
      statusCode: {
        404: function() {
          $('div#userFlash').delay(500).fadeIn('normal', function() {
            $(this).delay(5000).fadeOut()
          })
        }
      },

      // if username is valid, add it to `creds` object
      success: function (data) {
        creds['login'] = username;
        $('form#registration').hide()
        $('span#username').html(username)
        $('div#userSuccess').delay(200).fadeIn('normal', function() {
          $(this).show()
        })
      }
    })

    // Check is user has public keys linked to the github account
    $.ajax ({
      type: 'GET',
      url: url + '/keys',
      success: function (data) {

        // Users without keys will have 0-length data
        if (data.length > 0) {
          let keys = []
          for (let i = 0; i < data.length; i++) {
            keys.push(data[i]['key'])
          }
          creds['keys'] = keys;
          $('textarea#keySuccess').delay(500).fadeIn('normal', function() {
            $(this).text(creds.keys.join(',').replace(',', '\n', -1))
          })
          $('div#keyMsgSuccess').delay(500).fadeIn('normal', function() {
            $(this).show()
          })
          $('button#keyConfirmYes').click(function () {
            // send creds to backend
            $('div#keyMsgSuccess').delay(100).fadeIn('normal', function() {
              $(this).hide()
            })
            $('.service').delay(500).fadeIn('normal', function() {
              $(this).show()
            })
          })

          // Users without a key(s) will need to provide their own
        } else {
          $('div#keyMsgFailure').delay(500).fadeIn('normal', function() {
            $(this).show()
          })
          $('button#pk').click( function() {
            event.preventDefault()
            let key = $('textarea#pk').val()
            let result = verifyKey(key);
            if (result != null) {
              creds['keys'] = key
              $('div#keyMsgFailure').hide()
              $('div#keySuccess').delay(500).fadeIn('normal', function() {
                $(this).text('Key added successfully. Let\'s go!')
              })
            } else {
              $('div#badKeyFlash').delay(500).fadeIn('normal', function() {
                $(this).text('Invalid key. Please verify and try again.')
              })
            }
          })
        }
      }
    })
	})
})
