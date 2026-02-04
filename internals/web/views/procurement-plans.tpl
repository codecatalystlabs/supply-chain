{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Procurement Plans</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="procurementPlansTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Plan System Code</th>
          <th>Store Code</th>
          <th>Facility</th>
          <th>Financial Year</th>
          <th>Period Type</th>
          <th>Period Start</th>
          <th>Period End</th>
          <th>Items Count</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody id="procurementPlansBody">
        <tr>
          <td colspan="10" class="empty-state">
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
  loadProcurementPlans();
});

function loadProcurementPlans() {
  $.ajax({
    url: '/api/v1/procurement-plans',
    method: 'GET',
    success: function(data) {
      const tbody = $('#procurementPlansBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.plan_system_code || '-'}</td>
              <td>${item.store_code || '-'}</td>
              <td>${item.facility_name || item.facility_code || '-'}</td>
              <td>${item.financial_year || '-'}</td>
              <td>${item.plan_period_type || '-'}</td>
              <td>${item.plan_period_start ? new Date(item.plan_period_start).toLocaleDateString() : '-'}</td>
              <td>${item.plan_period_end ? new Date(item.plan_period_end).toLocaleDateString() : '-'}</td>
              <td>${item.items ? item.items.length : 0}</td>
              <td>${item.created_at ? new Date(item.created_at).toLocaleDateString() : '-'}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="10" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No procurement plans found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#procurementPlansBody').html(`
        <tr>
          <td colspan="10" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load procurement plans data');
    }
  });
}
</script>
{{end}}
