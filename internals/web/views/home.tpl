{{template "base.tpl" .}}

{{define "content"}}
<div class="row">
  <!-- Welcome Card -->
  <div class="col-12">
    <div class="card">
      <div class="card-body">
        <h4 class="card-title mb-2">Welcome back, {{if .Username}}{{.Username}}{{else}}User{{end}}!</h4>
        <p class="card-text mb-0">Supply Chain Management System Dashboard</p>
        {{if .Data}}{{with get .Data "user"}}{{if .FacilityName}}
        <p class="text-muted mb-0"><small><i class="icon ion-ios-location"></i> {{.FacilityName}}</small></p>
        {{end}}{{end}}{{end}}
      </div>
    </div>
  </div>
</div>

<div class="row mg-t-15">
  <!-- Facilities Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Facilities</h6>
            <h2 class="card-value" id="totalFacilities">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-location" style="font-size: 48px; color: #0066cc;"></i>
          </div>
        </div>
        <a href="/cp/facilities" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>

  <!-- Warehouses Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Warehouses</h6>
            <h2 class="card-value" id="totalWarehouses">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-home" style="font-size: 48px; color: #28a745;"></i>
          </div>
        </div>
        <a href="/cp/warehouses" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>

  <!-- Pharmacies Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Pharmacies</h6>
            <h2 class="card-value" id="totalPharmacies">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-medical" style="font-size: 48px; color: #dc3545;"></i>
          </div>
        </div>
        <a href="/cp/pharmacies" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>

  <!-- Procurement Plans Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Procurement Plans</h6>
            <h2 class="card-value" id="totalProcurementPlans">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-list" style="font-size: 48px; color: #ffc107;"></i>
          </div>
        </div>
        <a href="/cp/procurement-plans" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>
</div>

<div class="row mg-t-15">
  <!-- Purchase Orders Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Purchase Orders</h6>
            <h2 class="card-value" id="totalPurchaseOrders">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-cart" style="font-size: 48px; color: #17a2b8;"></i>
          </div>
        </div>
        <a href="/cp/purchase-orders" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>

  <!-- Stock on Hand Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Stock on Hand</h6>
            <h2 class="card-value" id="totalStockOnHand">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-box" style="font-size: 48px; color: #6f42c1;"></i>
          </div>
        </div>
        <a href="/cp/stock-on-hand" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>

  <!-- Stock Transfers Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Stock Transfers</h6>
            <h2 class="card-value" id="totalStockTransfers">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-swap" style="font-size: 48px; color: #20c997;"></i>
          </div>
        </div>
        <a href="/cp/stock-transfers" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>

  <!-- Pending Transfers Card -->
  <div class="col-sm-6 col-lg-3">
    <div class="card card-dashboard card-warning">
      <div class="card-body">
        <div class="d-flex justify-content-between align-items-center">
          <div>
            <h6 class="card-title">Pending Transfers</h6>
            <h2 class="card-value" id="pendingTransfers">0</h2>
          </div>
          <div class="card-icon">
            <i class="icon ion-ios-time" style="font-size: 48px; color: #ffc107;"></i>
          </div>
        </div>
        <a href="/cp/stock-transfers" class="card-link">View All <i class="fa fa-arrow-right"></i></a>
      </div>
    </div>
  </div>
</div>

<!-- Quick Actions -->
<div class="row mg-t-15">
  <div class="col-12">
    <div class="card">
      <div class="card-header">
        <h6 class="card-title mb-0">Quick Actions</h6>
      </div>
      <div class="card-body">
        <div class="row">
          <div class="col-md-3 col-sm-6 mg-b-10">
            <a href="/cp/procurement-plans" class="btn btn-outline-primary btn-block">
              <i class="icon ion-ios-add-circle-outline"></i> New Procurement Plan
            </a>
          </div>
          <div class="col-md-3 col-sm-6 mg-b-10">
            <a href="/cp/purchase-orders" class="btn btn-outline-success btn-block">
              <i class="icon ion-ios-cart-outline"></i> Create Purchase Order
            </a>
          </div>
          <div class="col-md-3 col-sm-6 mg-b-10">
            <a href="/cp/stock-transfers" class="btn btn-outline-info btn-block">
              <i class="icon ion-ios-swap"></i> New Stock Transfer
            </a>
          </div>
          <div class="col-md-3 col-sm-6 mg-b-10">
            <a href="/cp/stock-adjustments" class="btn btn-outline-warning btn-block">
              <i class="icon ion-ios-create-outline"></i> Stock Adjustment
            </a>
          </div>
        </div>
      </div>
    </div>
  </div>
</div>
{{end}}

{{define "extra_js"}}
<script>
$(document).ready(function() {
  loadDashboardStats();
});

function loadDashboardStats() {
  // Load stats from API endpoints
  Promise.all([
    $.get('/api/v1/facilities').then(data => ({facilities: data.length})),
    $.get('/api/v1/warehouses').then(data => ({warehouses: data.length})),
    $.get('/api/v1/pharmacies').then(data => ({pharmacies: data.length})),
    $.get('/api/v1/procurement-plans').then(data => ({procurementPlans: data.length})),
    $.get('/api/v1/purchase-orders').then(data => ({purchaseOrders: data.length})),
    $.get('/api/v1/stock/on-hand').then(data => ({stockOnHand: data.length})),
    $.get('/api/v1/stock/transfers').then(data => {
      const total = data.length;
      const pending = data.filter(t => t.status === 'pending').length;
      return {stockTransfers: total, pendingTransfers: pending};
    })
  ]).then(results => {
    const stats = Object.assign({}, ...results);
    
    $('#totalFacilities').text(stats.facilities || 0);
    $('#totalWarehouses').text(stats.warehouses || 0);
    $('#totalPharmacies').text(stats.pharmacies || 0);
    $('#totalProcurementPlans').text(stats.procurementPlans || 0);
    $('#totalPurchaseOrders').text(stats.purchaseOrders || 0);
    $('#totalStockOnHand').text(stats.stockOnHand || 0);
    $('#totalStockTransfers').text(stats.stockTransfers || 0);
    $('#pendingTransfers').text(stats.pendingTransfers || 0);
  }).catch(error => {
    console.error('Error loading dashboard stats:', error);
    toastr.error('Failed to load dashboard statistics');
  });
}
</script>

<style>
.card-dashboard {
  border-left: 4px solid #0066cc;
  transition: transform 0.2s ease, box-shadow 0.2s ease;
}

.card-dashboard:hover {
  transform: translateY(-2px);
  box-shadow: 0 4px 8px rgba(0,0,0,0.1);
}

.card-dashboard.card-warning {
  border-left-color: #ffc107;
}

.card-value {
  font-size: 28px;
  font-weight: 700;
  color: #333;
  margin: 8px 0;
  line-height: 1.2;
}

.card-title {
  font-size: 13px;
  font-weight: 600;
  color: #666;
  margin-bottom: 5px;
  text-transform: uppercase;
  letter-spacing: 0.5px;
}

.card-icon {
  opacity: 0.3;
}

.card-link {
  color: #0066cc;
  text-decoration: none;
  font-size: 14px;
  font-weight: 500;
  margin-top: 10px;
  display: inline-block;
}

.card-link:hover {
  text-decoration: underline;
}
</style>
{{end}}

