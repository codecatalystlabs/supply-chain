{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Product AMC</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
      <thead>
        <tr>
          <th>Product</th>
          <th>Facility</th>
          <th>AMC</th>
          <th>Period</th>
          <th>Last Updated</th>
          <th>Status</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.ProductName}}</td>
            <td>{{.FacilityName}}</td>
            <td>{{.AMC}}</td>
            <td>{{.Period}}</td>
            <td>{{.UpdatedAt}}</td>
            <td><span class="badge badge-success">Active</span></td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No product AMC data found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
