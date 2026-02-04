{{define "facilities.tpl"}}
{{template "base.tpl" .}}
{{end}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Health Facilities</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="facilitiesTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>Code</th>
          <th>Name</th>
          <th>Level of Care</th>
          <th>District</th>
          <th>Region</th>
          <th>Status</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody id="facilitiesBody">
        <tr>
          <td colspan="8" class="empty-state">
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
  loadFacilities();
});

function loadFacilities() {
  $.ajax({
    url: '/api/v1/facilities',
    method: 'GET',
    success: function(data) {
      const tbody = $('#facilitiesBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.facility_code || '-'}</td>
              <td>${item.facility_name || '-'}</td>
              <td>${item.level_of_care || '-'}</td>
              <td>${item.district || '-'}</td>
              <td>${item.region || '-'}</td>
              <td><span class="badge badge-${item.is_active ? 'success' : 'secondary'}">${item.is_active ? 'Active' : 'Inactive'}</span></td>
              <td>${item.created_at ? new Date(item.created_at).toLocaleDateString() : '-'}</td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="8" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No health facilities found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#facilitiesBody').html(`
        <tr>
          <td colspan="8" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load facilities data');
    }
  });
}
</script>
{{end}}
