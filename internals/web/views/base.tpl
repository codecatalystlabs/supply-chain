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
    
    <!-- Windows Table Styling -->
    <link rel="stylesheet" href="/cp/static/css/windows-table.css">

    <!-- Custom CSS for sidebar dropdown -->
    <style>
      /* Windows-like professional table styling - Global */
      .windows-table {
        border-collapse: collapse;
        width: 100%;
        font-size: 13px;
        background: #fff;
        box-shadow: 0 2px 4px rgba(0,0,0,0.15), 0 1px 2px rgba(0,0,0,0.1);
      }
      
      .windows-table thead {
        background: linear-gradient(to bottom, #f8f9fa 0%, #e9ecef 100%);
        border-bottom: 2px solid #dee2e6;
      }
      
      .windows-table thead th {
        padding: 6px 12px;
        font-weight: 600;
        font-size: 12px;
        color: #495057;
        text-transform: uppercase;
        letter-spacing: 0.3px;
        border-right: 1px solid #dee2e6;
        white-space: nowrap;
        text-align: left;
      }
      
      .windows-table thead th:last-child {
        border-right: none;
      }
      
      .windows-table tbody tr {
        border-bottom: 1px solid #e9ecef;
        transition: background-color 0.15s ease;
      }
      
      .windows-table tbody tr:hover {
        background-color: #e7f3ff;
      }
      
      .windows-table tbody tr:nth-child(even) {
        background-color: #fafbfc;
      }
      
      .windows-table tbody tr:nth-child(even):hover {
        background-color: #e7f3ff;
      }
      
      .windows-table tbody td {
        padding: 7px 12px;
        color: #212529;
        border-right: 1px solid #f0f0f0;
        vertical-align: middle;
      }
      
      .windows-table tbody td:last-child {
        border-right: none;
      }
      
      /* Badge styling */
      .status-badge {
        display: inline-block;
        padding: 3px 8px;
        font-size: 11px;
        font-weight: 600;
        border-radius: 3px;
        text-transform: uppercase;
        letter-spacing: 0.3px;
      }
      
      .status-badge.active {
        background-color: #d4edda;
        color: #155724;
        border: 1px solid #c3e6cb;
      }
      
      .status-badge.inactive {
        background-color: #f8d7da;
        color: #721c24;
        border: 1px solid #f5c6cb;
      }
      
      .status-badge.pending {
        background-color: #fff3cd;
        color: #856404;
        border: 1px solid #ffeaa7;
      }
      
      /* Card adjustments */
      .card {
        border: 1px solid #dee2e6;
        box-shadow: 0 2px 4px rgba(0,0,0,0.08);
      }
      
      .card-header {
        background: linear-gradient(to bottom, #ffffff 0%, #f8f9fa 100%);
        border-bottom: 1px solid #dee2e6;
        padding: 10px 15px;
      }
      
      .card-title {
        margin: 0;
        font-size: 14px;
        font-weight: 600;
        color: #495057;
      }
      
      .card-body {
        padding: 0;
      }
      
      .empty-state {
        text-align: center;
        padding: 40px 20px;
        color: #6c757d;
        font-size: 13px;
      }

      /* Multi-level sidebar navigation styles */
      .sidebar-nav-sub {
        display: none;
        background-color: rgba(0,0,0,0.03);
        padding: 2px 0;
        margin: 0;
      }
      
      .nav-sub-item {
        margin: 0;
      }
      
      .nav-sub-item:not(:last-child) {
        border-bottom: 1px solid rgba(0,0,0,0.05);
      }
      
      .nav-sub-link {
        padding: 8px 15px 8px 35px !important;
        font-size: 12px;
        display: block;
        transition: background-color 0.2s ease;
      }
      
      .nav-sub-link:hover {
        background-color: rgba(255,255,255,0.15);
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
              <span style="font-size: 14px; font-weight: 600; color: #333; font-family: 'Inter', sans-serif; letter-spacing: 0.3px;">SCM</span>
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
          
          <li class="sidebar-nav-item">
            <a href="/cp/dashboard" class="sidebar-nav-link"><i class="icon ion-ios-speedometer-outline"></i> Dashboard</a>
          </li>
          
          <!-- Facilities & Warehouses -->
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-location"></i> Facilities & Warehouses</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/facilities" class="nav-sub-link">Health Facilities</a></li>
              <li class="nav-sub-item"><a href="/cp/warehouses" class="nav-sub-link">Warehouses</a></li>
              <li class="nav-sub-item"><a href="/cp/pharmacies" class="nav-sub-link">Pharmacies</a></li>
            </ul>
          </li>

          <!-- Inventory & Stock -->
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-box"></i> Inventory & Stock</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/inventory" class="nav-sub-link">Inventory</a></li>
              <li class="nav-sub-item"><a href="/cp/pharmacy-stock" class="nav-sub-link">Pharmacy Stock</a></li>
              <li class="nav-sub-item"><a href="/cp/stock-on-hand" class="nav-sub-link">Stock on Hand</a></li>
              <li class="nav-sub-item"><a href="/cp/stock-dispensed" class="nav-sub-link">Stock Dispensed</a></li>
              <li class="nav-sub-item"><a href="/cp/stock-adjustments" class="nav-sub-link">Stock Adjustments</a></li>
              <li class="nav-sub-item"><a href="/cp/stock-returns" class="nav-sub-link">Stock Returns</a></li>
              <li class="nav-sub-item"><a href="/cp/stock-transfers" class="nav-sub-link">Stock Transfers</a></li>
            </ul>
          </li>

          <!-- Procurement & Orders -->
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-cart"></i> Procurement & Orders</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/procurement-plans" class="nav-sub-link">Procurement Plans</a></li>
              <li class="nav-sub-item"><a href="/cp/procurement-plan-import" class="nav-sub-link">Import Procurement Plan</a></li>
              <li class="nav-sub-item"><a href="/cp/purchase-orders" class="nav-sub-link">Purchase Orders</a></li>
              <li class="nav-sub-item"><a href="/cp/facility-orders" class="nav-sub-link">Facility Orders</a></li>
              <li class="nav-sub-item"><a href="/cp/warehouse-orders" class="nav-sub-link">Warehouse Orders</a></li>
              <li class="nav-sub-item"><a href="/cp/goods-receipt" class="nav-sub-link">Goods Receipts</a></li>
            </ul>
          </li>

          <!-- Patient Management -->
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-people"></i> Patient Management</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/patient-visits" class="nav-sub-link">Patient Visits</a></li>
              <li class="nav-sub-item"><a href="/cp/product-amc" class="nav-sub-link">Product AMC</a></li>
            </ul>
          </li>

          <!-- Reports -->
          <li class="sidebar-nav-item">
            <a href="/cp/reports" class="sidebar-nav-link"><i class="icon ion-stats-bars"></i> Reports</a>
          </li>

          <!-- Administration -->
          <li class="sidebar-nav-item with-sub">
            <a href="#" class="sidebar-nav-link"><i class="icon ion-ios-gear"></i> Administration</a>
            <ul class="nav sidebar-nav-sub">
              <li class="nav-sub-item"><a href="/cp/users" class="nav-sub-link">User Management</a></li>
              <li class="nav-sub-item"><a href="/cp/roles" class="nav-sub-link">Role Management</a></li>
            </ul>
          </li>

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

        <div class="slim-footer mg-t-0" style="padding: 3px 0 !important;">
          <div class="container-fluid" style="height: 30px !important;">
            <p style="margin: 0 !important; font-size: 10px; color: #6c757d; line-height: 1.2;"><strong>Supply Chain Management</strong> | All Rights Reserved. © 2026</p>
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
