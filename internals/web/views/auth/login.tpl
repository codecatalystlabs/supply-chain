<!DOCTYPE html>
<html lang="en">
  
<head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    {{/* <!-- Twitter -->
    <meta name="twitter:site" content="Login">
    <meta name="twitter:creator" content="Login">
    <meta name="twitter:card" content="summary_large_image">
    <meta name="twitter:title" content="Slim">
    <meta name="twitter:description" >
    <meta name="twitter:image" content="../../../slim/img/slim-social.html">

    <!-- Facebook -->
    <meta property="og:url" content="http://themepixels.me/slim">
    <meta property="og:title" content="Slim">
    <meta property="og:description">

    <meta property="og:image" content="../../../slim/img/slim-social.html">
    <meta property="og:image:secure_url" content="../../../slim/img/slim-social.html">
    <meta property="og:image:type" content="image/png">
    <meta property="og:image:width" content="1200">
    <meta property="og:image:height" content="600"> */}}

    <!-- Meta -->
    <meta name="description" content="Llincp.">

    <title>Login</title>

    <!-- Vendor css -->
    <link href="/cp/static/lib/font-awesome/css/font-awesome.css" rel="stylesheet">
    <link href="/cp/static/lib/Ionicons/css/ionicons.css" rel="stylesheet">

    <!-- Slim CSS -->
    <link rel="stylesheet" href="/cp/static/css/slim.css">
    
    <!-- Custom CSS for signin-box position -->
    <style>
      .signin-wrapper {
        display: flex !important;
        flex-direction: column !important;
        align-items: center !important;
        justify-content: center !important;
        min-height: 100vh !important;
        margin-top: -40px !important;
      }
      
      .logo-container {
        width: 100% !important;
        text-align: center !important;
        margin-bottom: 4px !important;
      }
      
      .slim-logo {
        margin: 0 !important;
        position: relative !important;
      }
      
      /* Password visibility toggle styling */
      .password-input-wrapper {
        position: relative;
        display: flex;
        align-items: center;
      }
      
      .password-input-wrapper input {
        flex: 1;
        padding-right: 38px;
      }
      
      .password-toggle-btn {
        position: absolute;
        right: 10px;
        background: none !important;
        border: none !important;
        cursor: pointer;
        color: #666;
        font-size: 16px;
        padding: 0 !important;
        display: flex;
        align-items: center;
        justify-content: center;
        width: 30px;
        height: 30px;
        z-index: 10;
        transition: color 0.2s ease;
      }
      
      .password-toggle-btn:hover {
        color: #333;
      }
      
      .password-toggle-btn:focus {
        outline: none;
        color: #333;
      }
      
      /* Enhanced login page styling */
      .login-header {
        text-align: center;
        margin-bottom: 30px;
      }
      
      .portal-title {
        font-size: 24px;
        font-weight: 600;
        color: #333;
        margin: 10px 0 10px 0;
        letter-spacing: 0.5px;
      }
      
      .portal-subtitle {
        font-size: 14px;
        color: #666;
        margin: 0;
        font-weight: 400;
      }
      
      .signin-box {
        box-shadow: 0 4px 15px rgba(0, 0, 0, 0.1) !important;
        border-radius: 10px !important;
        background: #fff;
      }
      
      .btn-signin {
        font-weight: 600;
        padding: 10px 15px !important;
        letter-spacing: 0.5px;
      }
    </style>

  </head>
  <body>

    <div class="signin-wrapper">
      <div class="login-header">
        <div class="logo-container">
          <img src="/cp/static/img/coat_fine.png" alt="SCM" style="max-width: 100px; height: 100px;">
        </div>
        <h3 class="portal-title">SCM</h3>
        <p class="portal-subtitle">Supply Chain Management Portal</p>
      </div>
      <div class="signin-box" style="padding:5px!important;">
        {{if .Errors.general}}
        <div style="background: #fed7d7; color: #c53030; padding: 12px 16px; border-radius: 8px; border-left: 4px solid #c53030; margin-bottom: 20px; font-size: 14px;">
          <i class="fa fa-exclamation-triangle" style="margin-right: 8px;"></i>
          {{.Errors.general}}
        </div>
        {{end}}
        
        <form method="POST" action="/cp/login">
          <div class="form-group" style="margin-top:2px;">
            <input type="text" class="form-control" name="username" placeholder="Username" required>
          </div><!-- form-group -->
          
          <div class="form-group">
            <div class="password-input-wrapper">
              <input type="password" class="form-control" id="passwordField" name="password" placeholder="Password" required>
              <button type="button" class="password-toggle-btn" id="passwordToggle">
                <i class="fa fa-eye"></i>
              </button>
            </div>
          </div><!-- form-group -->
          
          <button type="submit" class="btn btn-primary btn-block btn-signin" style="margin-bottom:2px !important;">
            Log In
          </button>
        </form>
      </div>
    </div><!-- signin-wrapper -->

    <script src="/cp/static/lib/jquery/js/jquery.js"></script>
    <script src="/cp/static/lib/popper.js/js/popper.js"></script>
    <script src="/cp/static/lib/bootstrap/js/bootstrap.js"></script>

    <script src="/cp/static/js/slim.js"></script>
    
    <script>
      // Password visibility toggle - vanilla JavaScript
      (function() {
        var toggleBtn = document.getElementById('passwordToggle');
        var passwordField = document.getElementById('passwordField');
        
        if (toggleBtn && passwordField) {
          toggleBtn.addEventListener('click', function(e) {
            e.preventDefault();
            e.stopPropagation();
            
            if (passwordField.type === 'password') {
              passwordField.type = 'text';
              toggleBtn.innerHTML = '<i class="fa fa-eye-slash"></i>';
            } else {
              passwordField.type = 'password';
              toggleBtn.innerHTML = '<i class="fa fa-eye"></i>';
            }
            
            passwordField.focus();
          });
        }
      })();
    </script>

  </body>
</html>
