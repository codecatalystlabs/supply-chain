{{define "sidebar_menu"}}
{{if .serviceResources}}
  {{range .serviceResources}}
    <li class="sidebar-nav-item">
      <a href="{{if .Domain}}{{.Domain}}{{else}}/services/{{.Id}}{{end}}" class="sidebar-nav-link">
        <i class="icon {{if .Icon}}{{.Icon}}{{else}}ion-ios-link-outline{{end}}"></i> 
        {{.Name}}
      </a>
    </li>
  {{end}}
{{end}}
{{end}}
