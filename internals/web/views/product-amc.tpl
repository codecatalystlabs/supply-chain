{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Product AMC</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="productAmcTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Product Code</th>
          <th>Product Name</th>
          <th>AMC Value</th>
          <th>Month</th>
          <th>Year</th>
          <th>Date</th>
          <th>Days Out of Stock</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="productAmcBody">
        <tr>
          <td colspan="11" class="empty-state">
            <i class="icon ion-ios-refresh" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Loading...
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
{{end}}

{{define "extra_js"}}
<script>
$(document).ready(function() {
  loadProductAMC();
});

function loadProductAMC() {
  $.ajax({
    url: '/api/v1/product-amc',
    method: 'GET',
    success: function(data) {
      const tbody = $('#productAmcBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.amc_system_code || '-'}</td>
              <td>${item.amc_facility_code || '-'}</td>
              <td>${item.amc_product_code || '-'}</td>
              <td>${item.amc_product_name || '-'}</td>
              <td><strong>${item.amc_value || 0}</strong></td>
              <td>${item.amc_month || '-'}</td>
              <td>${item.amc_year || '-'}</td>
              <td>${item.amc_date ? new Date(item.amc_date).toLocaleDateString() : '-'}</td>
              <td>${item.amc_days_out_stock || 0}</td>
              <td><span class="badge badge-${item.validation_status === 1 ? 'success' : 'warning'}">${item.validation_status === 1 ? 'Valid' : 'Pending'}</span></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="11" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No product AMC data found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#productAmcBody').html(`
        <tr>
          <td colspan="11" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load product AMC data');
    }
  });
}
</script>
{{end}}
