{{template "base.tpl" .}}

{{define "breadcrumb"}}
<ol class="breadcrumb slim-breadcrumb">
  <li class="breadcrumb-item"><a href="/">Home</a></li>
  <li class="breadcrumb-item active" aria-current="page">Dashboards</li>
</ol>
{{end}}

{{define "main_content"}}
<div class="section-wrapper">
  <div class="card">
    <div class="card-header">
      <h5 class="card-title mb-0">
        <i class="icon ion-stats-bars"></i> Dashboards
      </h5>
    </div>
    <div class="card-body">
      <!-- Dashboard Selection -->
      <div class="row mg-b-30">
        <div class="col-md-6">
          <div class="card card-dashboard-option">
            <div class="card-body">
              <div class="d-flex align-items-center">
                <div class="dashboard-icon bg-primary-light">
                  <i class="icon ion-map tx-primary tx-32"></i>
                </div>
                <div class="ml-3">
                  <h5 class="card-title mb-1">District Dashboard</h5>
                  <p class="tx-13 tx-gray-600 mb-0">View district-level statistics and coverage metrics</p>
                  <a href="/cp/dashboards/district" class="btn btn-sm btn-primary mt-3">
                    <i class="fa fa-arrow-right mr-1"></i>View Dashboard
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>

        <div class="col-md-6">
          <div class="card card-dashboard-option">
            <div class="card-body">
              <div class="d-flex align-items-center">
                <div class="dashboard-icon bg-success-light">
                  <i class="icon ion-ios-navigate tx-success tx-32"></i>
                </div>
                <div class="ml-3">
                  <h5 class="card-title mb-1">Admin Unit Drill-Down</h5>
                  <p class="tx-13 tx-gray-600 mb-0">Explore hierarchical admin units with detailed analytics</p>
                  <a href="/cp/dashboards/admin-unit" class="btn btn-sm btn-success mt-3">
                    <i class="fa fa-arrow-right mr-1"></i>View Dashboard
                  </a>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  .card-dashboard-option {
    border: 1px solid #e9ecef;
    transition: all 0.3s ease;
    cursor: pointer;
  }
  
  .card-dashboard-option:hover {
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    transform: translateY(-2px);
  }
  
  .dashboard-icon {
    width: 64px;
    height: 64px;
    border-radius: 8px;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-shrink: 0;
  }
  
  .bg-primary-light {
    background-color: rgba(31, 93, 234, 0.1);
  }
  
  .bg-success-light {
    background-color: rgba(35, 183, 141, 0.1);
  }
  
  .tx-primary {
    color: #1f5dea;
  }
  
  .tx-success {
    color: #23b78d;
  }
  
  .tx-32 {
    font-size: 32px;
  }
  
  .dashboard-container {
    margin-top: 20px;
  }
  
  .embed-responsive iframe {
    display: block;
  }
</style>
{{end}}
