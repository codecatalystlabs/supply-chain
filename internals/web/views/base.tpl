<!DOCTYPE html>
<html lang="en">
  <head>
    <!-- Required meta tags -->
    <meta charset="utf-8">
    <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no">

    <!-- Meta -->
    <meta name="description" content="LLin CP">
    <meta name="author" content="LLin CP">

    <title>{{.Title}} </title>

    <!-- vendor css -->
  <link href="/cp/static/lib/font-awesome/css/font-awesome.css" rel="stylesheet">
  <link href="/cp/static/lib/Ionicons/css/ionicons.css" rel="stylesheet">
  <!-- Toastr CSS -->
  <link href="/cp/static/lib/toastr/toastr.min.css" rel="stylesheet">
    
    <!-- Google Fonts - Inter -->
    <link href="/cp/static/css/google-font-inter.css" rel="stylesheet">

    <!-- Slim CSS -->
    <link rel="stylesheet" href="/cp/static/css/slim.css">

    <!-- Custom CSS for sidebar dropdown -->
    <style>
      /* Multi-level sidebar navigation styles */
      .sidebar-nav-sub {
        display: none;
        background-color: rgba(0,0,0,0.03);
      }
      
      .sidebar-nav-sub-sub {
        display: none;
        padding-left: 15px;
        background-color: rgba(0,0,0,0.05);
        border-radius: 4px;
        margin-top: 2px;
      }
      
      /* Show submenus when parent is open */
      .sidebar-nav-item.with-sub.open > .sidebar-nav-sub {
        display: block;
      }
      
      .nav-sub-item.with-sub .sidebar-nav-sub-sub {
        display: block;
      }
      
      /* Arrow indicators */
      .sidebar-nav-item.with-sub > .sidebar-nav-link::after,
      .nav-sub-item.with-sub > .nav-sub-link::after {
        content: "\f105";
        font-family: "FontAwesome";
        position: absolute;
        right: 10px;
        top: 50%;
        transform: translateY(-50%);
        transition: transform 0.3s ease;
        font-size: 12px;
      }
      
      .sidebar-nav-item.with-sub > .sidebar-nav-link,
      .nav-sub-item.with-sub > .nav-sub-link {
        position: relative;
      }
      
      .nav-sub-item.with-sub.open > .nav-sub-link:after {
        transform: translateY(-50%) rotate(90deg);
      }
      
      .sidebar-nav-sub-sub .nav-sub-item .nav-sub-link {
        padding-left: 25px;
        font-size: 12px;
        opacity: 0.9;
      }
      
      .sidebar-nav-sub-sub .nav-sub-item .nav-sub-link:hover {
        opacity: 1;
        background-color: rgba(0,0,0,0.1);
      }
      
      /* Enhanced caret styling */
      .sidebar-nav-link .caret {
        float: right;
        margin-top: 8px;
        margin-right: 10px;
        opacity: 0.6;
        transition: transform 0.3s ease;
      }
      
      .sidebar-nav-item.with-sub.open .caret i {
        transform: rotate(90deg);
      }
      
      /* Improved hover effects */
      .sidebar-nav-item .sidebar-nav-link:hover,
      .nav-sub-item .nav-sub-link:hover {
        background-color: rgba(255,255,255,0.1);
        transition: background-color 0.2s ease;
      }
      
      /* Icon spacing */
      .sidebar-nav-link i,
      .nav-sub-link i {
        width: 20px;
        text-align: center;
        margin-right: 8px;
      }
    </style>

    {{block "extra_css" .}}{{end}}
  </head>


  <body>
    <div class="slim-header with-sidebar" style="height:55px !important;">
      <div class="container-fluid">
        <div class="slim-header-left">
          <h2 class="slim-logo">
            <a href="/">
              <img src="/cp/static/img/coat_fine.png" alt="SCM Logo" style="height: 58px; margin-right: 4px; vertical-align: middle;padding-top:1px !important;">
              {{/* <span style="font-size: 16px; font-weight: 600; color: #333;">..<span>.</span></span> */}}
              <span style="font-size: 14px; font-weight: 600; color: #333; font-family: 'Inter', sans-serif; letter-spacing: 0.3px;">Supply Chain Management</span>
            </a>
          </h2>
          <a href="#" id="slimSidebarMenu" class="slim-sidebar-menu" 
            style="width:35px !important; height:35px !important; margin-top:4px !important;"
          >
            <span></span>
          </a>
        </div><!-- slim-header-left -->
        <div class="slim-header-right">
          <div class="dropdown dropdown-c">
            <a href="#" class="logged-user" data-toggle="dropdown">
              {{/* <img src="/static/img/img1.jpg" alt=""> */}}
              <span>{{if .Username}}{{.Username}}{{else}}User{{end}}</span>
              <i class="fa fa-angle-down"></i>
            </a>
            <div class="dropdown-menu dropdown-menu-right">
              <nav class="nav">
                <a href="/cp/profile" class="nav-link"><i class="icon ion-person"></i> View Profile</a>
                <a href="/cp/settings" class="nav-link"><i class="icon ion-ios-gear"></i> Account Settings</a>
                <a href="/cp/logout" class="nav-link"><i class="icon ion-forward"></i> Sign Out</a>
              </nav>
            </div><!-- dropdown-menu -->
          </div><!-- dropdown -->
        </div><!-- header-right -->
      </div><!-- container-fluid -->
    </div><!-- slim-header -->

    <div class="slim-body">
      <div class="slim-sidebar">
        <label class="sidebar-label">Navigation</label>

        <ul class="nav nav-sidebar">
          <li class="sidebar-nav-item">
            <a href="/cp/home" class="sidebar-nav-link"><i class="icon ion-ios-home-outline"></i> Home</a>
          </li>
          
          <!-- Dynamic Services Menu - Disabled for now -->
          {{/*
          {{if .services}}
            <li class="sidebar-nav-item with-sub">
              <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-gear-outline"></i>Services</a>
              <ul class="nav sidebar-nav-sub">
              {{range .services}}
                <li class="nav-sub-item">
                  <a href="{{if and .RedirectUrl .OpenInIframe}}/services/{{.Id}}{{else if .RedirectUrl}}{{.RedirectUrl}}{{else}}/services/{{.Id}}{{end}}" 
                     class="nav-sub-link"
                     {{if and .RedirectUrl (not .OpenInIframe)}}target="_blank"{{end}}>
                    <i class="icon {{if .Icon}}{{.Icon}}{{else}}ion-ios-link-outline{{end}}"></i> 
                    {{.Name}}
                    {{if and .RedirectUrl (not .OpenInIframe)}}<i class="fa fa-external-link ml-1" style="font-size: 10px; opacity: 0.7;"></i>{{end}}
                  </a>
                </li>
              {{end}}
              </ul>
            </li>
          {{end}}
          */}}
          
          {{/* <!-- Static Management Menu - Only visible to users with management permissions -->
          {{if or .canManageUsers .canManageRoles .canManagePermissions}}
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-people-outline"></i> User Management</a>
            <ul class="nav sidebar-nav-sub">
              {{if .canManageUsers}}
              <li class="nav-sub-item"><a href="/admin/users/" class="nav-sub-link">Manage Users</a></li>
              {{end}}
              {{if .canManageRoles}}
              <li class="nav-sub-item"><a href="/admin/roles/" class="nav-sub-link">Manage Roles</a></li>
              {{end}}
              {{if .canManagePermissions}}
              <li class="nav-sub-item"><a href="/admin/permissions/" class="nav-sub-link">Manage Permissions</a></li>
              {{end}}
            </ul>
          </li>
          {{end}} */}}
          
          {{/* {{if .canManageServices}}
          <li class="sidebar-nav-item">
            <a href="/cp/services" class="sidebar-nav-link"><i class="icon ion-ios-cog-outline"></i> Service Management</a>
          </li>
          {{end}} */}}

          {{if .IsAdmin}}
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-people-outline"></i> User Management</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/admin/users/" class="nav-sub-link">Manage Users</a></li>
              <li class="nav-sub-item"><a href="/cp/admin/roles/" class="nav-sub-link">Manage Roles</a></li>
              <li class="nav-sub-item"><a href="/cp/admin/permissions/" class="nav-sub-link">Manage Permissions</a></li>
            </ul>
          </li>
          {{end}}
          
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-pulse"></i> Dashboards</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/dashboards/district" class="nav-sub-link">District</a></li>
              <li class="nav-sub-item"><a href="/cp/dashboards/admin-unit" class="nav-sub-link">Administrative Units</a></li>
            </ul>
          </li>

          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-stats-bars"></i> Reports</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/reports/vht" class="nav-sub-link">VHT</a></li>
            </ul>
          </li>

          {{if and .IsAdmin .CanExploreData}}
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-grid"></i> Data Exploration</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/admin/data-cleaning/duplicates" class="nav-sub-link">Flag Duplicates</a></li>
              <li class="nav-sub-item"><a href="/cp/admin/data-cleaning/training-mode" class="nav-sub-link">Misplaced Data</a></li>
              <li class="nav-sub-item"><a href="/cp/admin/data-cleaning/vht-submissions" class="nav-sub-link">VHT Submissions</a></li>
            </ul>
          </li>
          {{end}}

          {{/* <li class="sidebar-nav-item">
            <a href="/cp/reports" class="sidebar-nav-link"><i class="icon ion-stats-bars"></i> Reports</a>
          </li> */}}

          {{block "extra_menu" .}}{{end}}
        </ul>
      </div><!-- slim-sidebar -->

      <div class="slim-mainpanel">
        <div class="container" style="padding:10px !important;">
          <div class="slim-pageheader" style="padding:5px 0 !important;">
            {{block "breadcrumb" .}}
            <ol class="breadcrumb slim-breadcrumb">
              <li class="breadcrumb-item"><a href="/">Home</a></li>
              <li class="breadcrumb-item active" aria-current="page">{{.Title}}</li>
            </ol>
            {{end}}
          </div>

          {{block "content" .}}
          <div class="card">
            <div class="card-body">
              <p class="text-muted">No content specified</p>
            </div>
          </div>
          {{end}}


          {{/* {{block "main_content" .}}
          <div class="card">
            <div class="card-body">
              <h5 class="card-title">{{.title}}</h5>
              <p class="card-text">Main content goes here.</p>
            </div>
          </div>
          {{end}} */}}

        </div><!-- container -->

        <div class="slim-footer mg-t-0">
          <div class="container-fluid">
            <p><strong>Supply Chain Management System</strong> | All Rights Reserved. © 2026</p>
          </div><!-- container-fluid -->
        </div><!-- slim-footer -->
      </div><!-- slim-mainpanel -->
    </div><!-- slim-body -->

  <script src="/cp/static/lib/jquery/js/jquery.js"></script>
  <script src="/cp/static/lib/popper.js/js/popper.js"></script>
  <script src="/cp/static/lib/bootstrap/js/bootstrap.js"></script>
  <script src="/cp/static/lib/jquery.cookie/js/jquery.cookie.js"></script>
  <script src="/cp/static/lib/perfect-scrollbar/js/perfect-scrollbar.jquery.min.js"></script>

  <!-- Toastr JS -->
  <script src="/cp/static/lib/toastr/toastr.min.js"></script>

  <script src="/cp/static/js/slim.js"></script>

  {{block "extra_js" .}}{{end}}
  </body>
</html>
