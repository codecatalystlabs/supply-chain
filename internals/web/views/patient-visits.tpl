{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Patient Visits</h6>
  </div>
  <div class="card-body">
    <table class="windows-table">
      <thead>
        <tr>
          <th>Visit ID</th>
          <th>Patient</th>
          <th>Facility</th>
          <th>Visit Date</th>
          <th>Type</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.VisitID}}</td>
            <td>{{.PatientName}}</td>
            <td>{{.FacilityName}}</td>
            <td>{{.VisitDate}}</td>
            <td>{{.VisitType}}</td>
            <td><span class="badge badge-info">{{.Status}}</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr>
            <td colspan="6" class="empty-state">
              <i class="icon ion-ios-information-outline" style="font-size: 24px; display: block; margin-bottom: 8px;"></i>
              No patient visits found
            </td>
          </tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
