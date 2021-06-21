/*
 * ETSI ISG CIM / NGSI-LD API
 *
 * This OAS file describes the NGSI-LD API defined by the ETSI ISG CIM group. This Cross-domain Context Information Management API allows to provide, consume and subscribe to context information in multiple scenarios and involving multiple stakeholders
 *
 * API version: latest
 * Contact: NGSI-LD@etsi.org
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */
package models
import (
	"time"
)

type NotificationParams struct {

	Attributes []string `json:"attributes,omitempty"`

	Format string `json:"format,omitempty"`

	Endpoint *Endpoint `json:"endpoint"`

	Status string `json:"status,omitempty"`

	TimesSent float64 `json:"timesSent,omitempty"`

	LastNotification time.Time `json:"lastNotification,omitempty"`

	LastFailure time.Time `json:"lastFailure,omitempty"`

	LastSuccess time.Time `json:"lastSuccess,omitempty"`
}
