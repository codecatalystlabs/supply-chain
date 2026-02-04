{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Pharmacies</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="pharmaciesTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Code</th>
          <th>Name</th>
          <th>Type</th>
          <th>Facility</th>
          <th>Status</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody id="pharmaciesBody">
        <tr>
          <td colspan="7" class="empty-state">
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
  loadPharmacies();
});

function loadPharmacies() {
  $.ajax({
    url: '/api/v1/pharmacies',
    method: 'GET',
    success: function(data) {
      const tbody = $('#pharmaciesBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.pharmacy_code || '-'}</td>
              <td>${item.pharmacy_name || '-'}</td>
              <td>${item.pharmacy_type || '-'}</td>
              <td>${item.facility ? item.facility.facility_name : '-'}</td>
              <td><span class="badge badge-${item.is_active ? 'success' : 'secondary'}">${item.is_active ? 'Active' : 'Inactive'}</span></td>
              <td>${item.created_at ? new Date(item.created_at).toLocaleDateString() : '-'}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="7" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No pharmacies found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#pharmaciesBody').html(`
        <tr>
          <td colspan="7" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load pharmacies data');
    }
  });
}
</script>
{{end}}
