{{template "base.tpl" .}}

{{define "breadcrumb"}}
<ol class="breadcrumb slim-breadcrumb">
  <li class="breadcrumb-item"><a href="/">Home</a></li>
  <li class="breadcrumb-item active" aria-current="page">Reports</li>
</ol>
{{end}}

{{define "main_content"}}
<div class="section-wrapper">
  <div class="card">
    <div class="card-header">
      <h5 class="card-title mb-0">
        <i class="icon ion-stats-bars"></i> Reports
      </h5>
    </div>
    <div class="card-body">
      <p class="text-muted mb-3">
        Embed your Elasticsearch reports below. Replace the placeholder URLs with your actual Elasticsearch/Kibana report URLs.
      </p>
      
      <!-- Multiple Report Sections -->
      <div class="row">
        <!-- Report 1 -->
        <div class="col-md-12 mb-4">
          <div class="card">
            <div class="card-header bg-primary text-white">
              <h6 class="mb-0">Report 1 - Sample Report</h6>
            </div>
            <div class="card-body p-0">
              <div class="embed-responsive" style="position: relative; width: 100%; height: 600px; border: 1px solid #dee2e6;">
                <iframe 
                  src="about:blank"
                  data-placeholder-url="YOUR_ELASTICSEARCH_REPORT_1_URL"
                  style="width: 100%; height: 100%; border: 0;"
                  frameborder="0"
                  allowfullscreen>
                </iframe>
                <div style="position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); text-align: center; color: #6c757d;">
                  <i class="icon ion-stats-bars" style="font-size: 48px; opacity: 0.3;"></i>
                  <p class="mt-3">Report 1 iframe placeholder</p>
                  <p class="small">Update the iframe src to embed your report</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Report 2 -->
        <div class="col-md-12 mb-4">
          <div class="card">
            <div class="card-header bg-success text-white">
              <h6 class="mb-0">Report 2 - Analytics Report</h6>
            </div>
            <div class="card-body p-0">
              <div class="embed-responsive" style="position: relative; width: 100%; height: 600px; border: 1px solid #dee2e6;">
                <iframe 
                  src="about:blank"
                  data-placeholder-url="YOUR_ELASTICSEARCH_REPORT_2_URL"
                  style="width: 100%; height: 100%; border: 0;"
                  frameborder="0"
                  allowfullscreen>
                </iframe>
                <div style="position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); text-align: center; color: #6c757d;">
                  <i class="icon ion-stats-bars" style="font-size: 48px; opacity: 0.3;"></i>
                  <p class="mt-3">Report 2 iframe placeholder</p>
                  <p class="small">Update the iframe src to embed your report</p>
                </div>
              </div>
            </div>
          </div>
        </div>

        <!-- Report 3 -->
        <div class="col-md-12 mb-4">
          <div class="card">
            <div class="card-header bg-info text-white">
              <h6 class="mb-0">Report 3 - Performance Metrics</h6>
            </div>
            <div class="card-body p-0">
              <div class="embed-responsive" style="position: relative; width: 100%; height: 600px; border: 1px solid #dee2e6;">
                <iframe 
                  src="about:blank"
                  data-placeholder-url="YOUR_ELASTICSEARCH_REPORT_3_URL"
                  style="width: 100%; height: 100%; border: 0;"
                  frameborder="0"
                  allowfullscreen>
                </iframe>
                <div style="position: absolute; top: 50%; left: 50%; transform: translate(-50%, -50%); text-align: center; color: #6c757d;">
                  <i class="icon ion-stats-bars" style="font-size: 48px; opacity: 0.3;"></i>
                  <p class="mt-3">Report 3 iframe placeholder</p>
                  <p class="small">Update the iframe src to embed your report</p>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>

      <!-- Instructions -->
      <div class="card bg-light">
        <div class="card-body">
          <h6 class="card-title">How to embed your Elasticsearch reports:</h6>
          <ol class="mb-0">
            <li>Open your Elasticsearch/Kibana visualization or saved search</li>
            <li>Click the "Share" button in the top menu</li>
            <li>Select "Embed code" or copy the visualization URL</li>
            <li>Replace the <code>src="about:blank"</code> in the desired iframe above with your report URL</li>
            <li>Configure appropriate CORS and security settings in your Elasticsearch/Kibana instance</li>
            <li>Add authentication if required (consider using reverse proxy for secure embedding)</li>
          </ol>
        </div>
      </div>
    </div>
  </div>
</div>

<style>
  .embed-responsive iframe {
    display: block;
  }
  
  .card-header h6 {
    font-weight: 500;
  }
</style>
{{end}}
