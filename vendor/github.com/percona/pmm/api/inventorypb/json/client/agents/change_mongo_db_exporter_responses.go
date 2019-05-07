// Code generated by go-swagger; DO NOT EDIT.

package agents

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// ChangeMongoDBExporterReader is a Reader for the ChangeMongoDBExporter structure.
type ChangeMongoDBExporterReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ChangeMongoDBExporterReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewChangeMongoDBExporterOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewChangeMongoDBExporterDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewChangeMongoDBExporterOK creates a ChangeMongoDBExporterOK with default headers values
func NewChangeMongoDBExporterOK() *ChangeMongoDBExporterOK {
	return &ChangeMongoDBExporterOK{}
}

/*ChangeMongoDBExporterOK handles this case with default header values.

A successful response.
*/
type ChangeMongoDBExporterOK struct {
	Payload *ChangeMongoDBExporterOKBody
}

func (o *ChangeMongoDBExporterOK) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/ChangeMongoDBExporter][%d] changeMongoDbExporterOk  %+v", 200, o.Payload)
}

func (o *ChangeMongoDBExporterOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeMongoDBExporterOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewChangeMongoDBExporterDefault creates a ChangeMongoDBExporterDefault with default headers values
func NewChangeMongoDBExporterDefault(code int) *ChangeMongoDBExporterDefault {
	return &ChangeMongoDBExporterDefault{
		_statusCode: code,
	}
}

/*ChangeMongoDBExporterDefault handles this case with default header values.

An error response.
*/
type ChangeMongoDBExporterDefault struct {
	_statusCode int

	Payload *ChangeMongoDBExporterDefaultBody
}

// Code gets the status code for the change mongo DB exporter default response
func (o *ChangeMongoDBExporterDefault) Code() int {
	return o._statusCode
}

func (o *ChangeMongoDBExporterDefault) Error() string {
	return fmt.Sprintf("[POST /v0/inventory/Agents/ChangeMongoDBExporter][%d] ChangeMongoDBExporter default  %+v", o._statusCode, o.Payload)
}

func (o *ChangeMongoDBExporterDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(ChangeMongoDBExporterDefaultBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*ChangeMongoDBExporterBody change mongo DB exporter body
swagger:model ChangeMongoDBExporterBody
*/
type ChangeMongoDBExporterBody struct {

	// agent id
	AgentID string `json:"agent_id,omitempty"`

	// Replace all custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// disabled
	Disabled bool `json:"disabled,omitempty"`

	// enabled
	Enabled bool `json:"enabled,omitempty"`

	// Remove all custom user-assigned labels.
	RemoveCustomLabels bool `json:"remove_custom_labels,omitempty"`
}

// Validate validates this change mongo DB exporter body
func (o *ChangeMongoDBExporterBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterBody) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterDefaultBody ErrorResponse is a message returned on HTTP error.
swagger:model ChangeMongoDBExporterDefaultBody
*/
type ChangeMongoDBExporterDefaultBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	// error
	Error string `json:"error,omitempty"`

	// message
	Message string `json:"message,omitempty"`
}

// Validate validates this change mongo DB exporter default body
func (o *ChangeMongoDBExporterDefaultBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterDefaultBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterDefaultBody) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterDefaultBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterOKBody change mongo DB exporter OK body
swagger:model ChangeMongoDBExporterOKBody
*/
type ChangeMongoDBExporterOKBody struct {

	// mongodb exporter
	MongodbExporter *ChangeMongoDBExporterOKBodyMongodbExporter `json:"mongodb_exporter,omitempty"`
}

// Validate validates this change mongo DB exporter OK body
func (o *ChangeMongoDBExporterOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateMongodbExporter(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *ChangeMongoDBExporterOKBody) validateMongodbExporter(formats strfmt.Registry) error {

	if swag.IsZero(o.MongodbExporter) { // not required
		return nil
	}

	if o.MongodbExporter != nil {
		if err := o.MongodbExporter.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("changeMongoDbExporterOk" + "." + "mongodb_exporter")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBody) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*ChangeMongoDBExporterOKBodyMongodbExporter MongoDBExporter runs on Generic or Container Node and exposes MongoDB Service metrics.
swagger:model ChangeMongoDBExporterOKBodyMongodbExporter
*/
type ChangeMongoDBExporterOKBodyMongodbExporter struct {

	// Unique randomly generated instance identifier.
	AgentID string `json:"agent_id,omitempty"`

	// Custom user-assigned labels.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Desired Agent status: enabled (false) or disabled (true).
	Disabled bool `json:"disabled,omitempty"`

	// Listen port for scraping metrics.
	ListenPort int64 `json:"listen_port,omitempty"`

	// MongoDB password for scraping metrics.
	Password string `json:"password,omitempty"`

	// The pmm-agent identifier which runs this instance.
	PMMAgentID string `json:"pmm_agent_id,omitempty"`

	// Service identifier.
	ServiceID string `json:"service_id,omitempty"`

	// AgentStatus represents actual Agent status.
	// Enum: [AGENT_STATUS_INVALID STARTING RUNNING WAITING STOPPING DONE]
	Status *string `json:"status,omitempty"`

	// MongoDB username for scraping metrics.
	Username string `json:"username,omitempty"`
}

// Validate validates this change mongo DB exporter OK body mongodb exporter
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["AGENT_STATUS_INVALID","STARTING","RUNNING","WAITING","STOPPING","DONE"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum = append(changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum, v)
	}
}

const (

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID captures enum value "AGENT_STATUS_INVALID"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusAGENTSTATUSINVALID string = "AGENT_STATUS_INVALID"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusSTARTING captures enum value "STARTING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusSTARTING string = "STARTING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusRUNNING captures enum value "RUNNING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusRUNNING string = "RUNNING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusWAITING captures enum value "WAITING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusWAITING string = "WAITING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusSTOPPING captures enum value "STOPPING"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusSTOPPING string = "STOPPING"

	// ChangeMongoDBExporterOKBodyMongodbExporterStatusDONE captures enum value "DONE"
	ChangeMongoDBExporterOKBodyMongodbExporterStatusDONE string = "DONE"
)

// prop value enum
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, changeMongoDbExporterOkBodyMongodbExporterTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (o *ChangeMongoDBExporterOKBodyMongodbExporter) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(o.Status) { // not required
		return nil
	}

	// value enum
	if err := o.validateStatusEnum("changeMongoDbExporterOk"+"."+"mongodb_exporter"+"."+"status", "body", *o.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *ChangeMongoDBExporterOKBodyMongodbExporter) UnmarshalBinary(b []byte) error {
	var res ChangeMongoDBExporterOKBodyMongodbExporter
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
