package defaults

var SampleConfig = []byte(`config: 
  export_directory: /home/kasun/.wso2apimcli/exported
  http_request_timeout: 10000
environments: 
  sample-env1: 
    admin_endpoint: "https://localhost:9443/api/am/admin/v0.14"
    api_import_export_endpoint: "https://localhost/api-import-export"
    api_list_endpoint: "https://localhost/publisher/apis"
    api_manager_endpoint: "https://localhost/apim"
    application_list_endpoint: "https://localhost:9443/api/am/admin/v0.14/applications"
    registration_endpoint: "https://localhost/register"
    token_endpoint: "https://localhost/token"
  sample-env2: 
    admin_endpoint: ""
    api_import_export_endpoint: ""
    api_list_endpoint: ""
    api_manager_endpoint: "https://localhost/apim"
    application_list_endpoint: ""
    registration_endpoint: "https://localhost/register"
    token_endpoint: "https://localhost/token"
`)
