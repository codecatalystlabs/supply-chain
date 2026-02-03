{{template "base.tpl" .}}

{{define "content"}}
<div class="card">
  <div class="card-header">
    <h6 class="card-title">Procurement Plans</h6>
  </div>
  <div class="card-body">
    <table class="table table-hover table-striped">
      <thead>
        <tr>
          <th>Plan ID</th>
          <th>Period</th>
          <th>Status</th>
          <th>Total Budget</th>
          <th>Items</th>
          <th>Created Date</th>
        </tr>
      </thead>
      <tbody>
        {{if .Data}}
          {{range .Data}}
          <tr>
            <td>{{.PlanID}}</td>
            <td>{{.Period}}</td>
            <td><span class="badge badge-info">{{.Status}}</span></td>
            <td>{{.TotalBudget}}</td>
            <td>{{.ItemCount}}</td>
            <td>{{.CreatedDate}}</td>
          </tr>
          {{end}}
        {{else}}
          <tr><td colspan="6" class="text-center text-muted">No procurement plans found</td></tr>
        {{end}}
      </tbody>
    </table>
  </div>
</div>
{{end}}
