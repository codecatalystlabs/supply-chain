{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Pharmacies</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
      <thead>
        <tr>
          <th>ID</th>
          <th>Name</th>
          <th>Facility</th>
          <th>Head Pharmacist</th>
          <th>Status</th>
          <th>Created</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ID}}</td>
            <td>{{.Name}}</td>
            <td>{{.FacilityName}}</td>
            <td>{{.HeadPharmacist}}</td>
            <td><span class="badge badge-success">Active</span></td>
            <td>{{.CreatedAt}}</td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No pharmacies found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
