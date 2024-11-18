package constants

const ModuleName = "gateway"
const HttpServerPort = "8080"
const HttpGet = "GET"
const HttpPost = "POST"
const ApiBasePath = "/api/v1"
const GatewayServicePath = "/gateway"
const CustomerServicePath = "/customerservice"
const ProductServicePath = "/productservice"
const OrderServicePath = "/oms"

// Service Addresses
const (
	CustomerServiceAddressEnv = "CUSTOMER_SERVICE_ADDRESS"
	ProductServiceAddressEnv  = "PRODUCT_SERVICE_ADDRESS"
	OrderServiceAddressEnv    = "ORDER_SERVICE_ADDRESS"
	CustomerServiceAddress    = "127.0.0.1:8082"
	ProductServiceAddress     = "127.0.0.1:8083"
	OrderServiceAddress       = "127.0.0.1:8084"
)

// Otel
const (
	OtelEnableEnv       = "OTEL_ENABLED"
	OtelCollectorEnv    = "OTEL_COLLECTOR_URL"
	OtelInsecureModeEnv = "OTEL_INSECURE_MODE"
)
