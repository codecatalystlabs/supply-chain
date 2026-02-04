{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Patient Visits</h6>
  </div>
  <div class="card-body">
    <table class="windows-table" id="patientVisitsTable">
      <thead>
        <tr>
          <th>ID</th>
          <th>System Code</th>
          <th>Facility Code</th>
          <th>Patient Hash</th>
          <th>Visit Date</th>
          <th>Visit Type</th>
          <th>Timestamp</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody id="patientVisitsBody">
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
  loadPatientVisits();
});

function loadPatientVisits() {
  $.ajax({
    url: '/api/v1/patient-visit',
    method: 'GET',
    success: function(data) {
      const tbody = $('#patientVisitsBody');
      tbody.empty();
      
      if (data && data.length > 0) {
        data.forEach(function(item) {
          const row = `
            <tr>
              <td>${item.id}</td>
              <td>${item.pvt_system_code || '-'}</td>
              <td>${item.pvt_facility_code || '-'}</td>
              <td>${item.pvt_patient_hash ? item.pvt_patient_hash.substring(0, 8) + '...' : '-'}</td>
              <td>${item.pvt_visit_date ? new Date(item.pvt_visit_date).toLocaleDateString() : '-'}</td>
              <td>${item.pvt_visit_type || '-'}</td>
              <td>${item.pvt_timestamp ? new Date(item.pvt_timestamp).toLocaleString() : '-'}</td>
              <td><span class="badge badge-${item.validation_status === 1 ? 'success' : 'warning'}">${item.validation_status === 1 ? 'Valid' : 'Pending'}</span></td>
            </tr>
          `;
          tbody.append(row);
        });
      } else {
        tbody.append(`
          <tr>
            <td colspan="8" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No patient visits found
            </td>
          </tr>
        `);
      }
    },
    error: function(xhr, status, error) {
      $('#patientVisitsBody').html(`
        <tr>
          <td colspan="8" class="empty-state">
            <i class="icon ion-ios-close-circle-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
            Error loading data: ${error}
          </td>
        </tr>
      `);
      toastr.error('Failed to load patient visits data');
    }
  });
}
</script>
{{end}}
