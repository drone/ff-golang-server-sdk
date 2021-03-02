// Package rest provides primitives to interact the openapi HTTP API.
//
// Code generated by github.com/deepmap/oapi-codegen DO NOT EDIT.
package rest

// AuthenticationRequest defines model for AuthenticationRequest.
type AuthenticationRequest struct {
	ApiKey string `json:"apiKey"`
}

// AuthenticationResponse defines model for AuthenticationResponse.
type AuthenticationResponse struct {
	AuthToken string `json:"authToken"`
}

// Clause defines model for Clause.
type Clause struct {
	Attribute string   `json:"attribute"`
	Id        string   `json:"id"`
	Negate    bool     `json:"negate"`
	Op        string   `json:"op"`
	Values    []string `json:"values"`
}

// Distribution defines model for Distribution.
type Distribution struct {
	BucketBy   string              `json:"bucketBy"`
	Variations []WeightedVariation `json:"variations"`
}

// Error defines model for Error.
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Evaluation defines model for Evaluation.
type Evaluation struct {
	Flag  string      `json:"flag"`
	Value interface{} `json:"value"`
}

// Evaluations defines model for Evaluations.
type Evaluations []Evaluation

// FeatureConfig defines model for FeatureConfig.
type FeatureConfig struct {
	DefaultServe         Serve           `json:"defaultServe"`
	Environment          string          `json:"environment"`
	Feature              string          `json:"feature"`
	Kind                 string          `json:"kind"`
	OffVariation         string          `json:"offVariation"`
	Prerequisites        *[]Prerequisite `json:"prerequisites,omitempty"`
	Project              string          `json:"project"`
	Rules                *[]ServingRule  `json:"rules,omitempty"`
	State                FeatureState    `json:"state"`
	VariationToTargetMap *[]VariationMap `json:"variationToTargetMap,omitempty"`
	Variations           []Variation     `json:"variations"`
	Version              *int64          `json:"version,omitempty"`
}

// FeatureState defines model for FeatureState.
type FeatureState string

// List of FeatureState
const (
	FeatureState_off FeatureState = "off"
	FeatureState_on  FeatureState = "on"
)

// Pagination defines model for Pagination.
type Pagination struct {
	ItemCount int  `json:"itemCount"`
	PageCount int  `json:"pageCount"`
	PageIndex int  `json:"pageIndex"`
	PageSize  int  `json:"pageSize"`
	Version   *int `json:"version,omitempty"`
}

// Prerequisite defines model for Prerequisite.
type Prerequisite struct {
	Feature    string   `json:"feature"`
	Variations []string `json:"variations"`
}

// Segment defines model for Segment.
type Segment struct {
	CreatedAt   *int64    `json:"createdAt,omitempty"`
	Environment *string   `json:"environment,omitempty"`
	Excluded    *[]string `json:"excluded,omitempty"`

	// Unique identifier for the segment.
	Identifier string    `json:"identifier"`
	Included   *[]string `json:"included,omitempty"`
	ModifiedAt *int64    `json:"modifiedAt,omitempty"`

	// Name of the segment.
	Name string `json:"name"`

	// An array of rules that can cause a user to be included in this segment.
	Rules   *[]Clause `json:"rules,omitempty"`
	Tags    *[]Tag    `json:"tags,omitempty"`
	Version *int64    `json:"version,omitempty"`
}

// Serve defines model for Serve.
type Serve struct {
	Distribution *Distribution `json:"distribution,omitempty"`
	Variation    *string       `json:"variation,omitempty"`
}

// ServingRule defines model for ServingRule.
type ServingRule struct {
	Clauses  []Clause `json:"clauses"`
	Priority int      `json:"priority"`
	RuleId   string   `json:"ruleId"`
	Serve    Serve    `json:"serve"`
}

// Tag defines model for Tag.
type Tag struct {
	Name  string  `json:"name"`
	Value *string `json:"value,omitempty"`
}

// Variation defines model for Variation.
type Variation struct {
	Description *string `json:"description,omitempty"`
	Identifier  string  `json:"identifier"`
	Name        *string `json:"name,omitempty"`
	Value       string  `json:"value"`
}

// VariationMap defines model for VariationMap.
type VariationMap struct {
	TargetSegments *[]string `json:"targetSegments,omitempty"`
	Targets        *[]string `json:"targets,omitempty"`
	Variation      string    `json:"variation"`
}

// WeightedVariation defines model for WeightedVariation.
type WeightedVariation struct {
	Variation string `json:"variation"`
	Weight    int    `json:"weight"`
}

// InternalServerError defines model for InternalServerError.
type InternalServerError Error

// NotFound defines model for NotFound.
type NotFound Error

// Unauthenticated defines model for Unauthenticated.
type Unauthenticated Error

// Unauthorized defines model for Unauthorized.
type Unauthorized Error

// AuthenticateJSONBody defines parameters for Authenticate.
type AuthenticateJSONBody AuthenticationRequest

// AuthenticateRequestBody defines body for Authenticate for application/json ContentType.
type AuthenticateJSONRequestBody AuthenticateJSONBody

