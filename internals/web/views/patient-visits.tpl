{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Patient Visits</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
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
          <tr><td colspan="6" class="text-center text-muted">No patient visits found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
