// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"reflect"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func regionSslPolicyCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	profile := diff.Get("profile")
	customFeaturesCount := diff.Get("custom_features.#")

	// Validate that policy configs aren't incompatible during all phases
	// CUSTOM profile demands non-zero custom_features, and other profiles (i.e., not CUSTOM) demand zero custom_features
	if diff.HasChange("profile") || diff.HasChange("custom_features") {
		if profile.(string) == "CUSTOM" && customFeaturesCount.(int) == 0 {
			return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but no custom_features are set.", diff.Get("name"), profile.(string))
		} else if profile.(string) != "CUSTOM" && customFeaturesCount.(int) != 0 {
			return fmt.Errorf("Error in SSL Policy %s: the profile is set to %s but using custom_features requires the profile to be CUSTOM.", diff.Get("name"), profile.(string))
		}
		return nil
	}
	return nil
}

func resourceComputeRegionSslPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceComputeRegionSslPolicyCreate,
		Read:   resourceComputeRegionSslPolicyRead,
		Update: resourceComputeRegionSslPolicyUpdate,
		Delete: resourceComputeRegionSslPolicyDelete,

		Importer: &schema.ResourceImporter{
			State: resourceComputeRegionSslPolicyImport,
		},

		Timeouts: &schema.ResourceTimeout{
			Create: schema.DefaultTimeout(20 * time.Minute),
			Update: schema.DefaultTimeout(20 * time.Minute),
			Delete: schema.DefaultTimeout(20 * time.Minute),
		},

		CustomizeDiff: regionSslPolicyCustomizeDiff,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
				Description: `Name of the resource. Provided by the client when the resource is
created. The name must be 1-63 characters long, and comply with
RFC1035. Specifically, the name must be 1-63 characters long and match
the regular expression '[a-z]([-a-z0-9]*[a-z0-9])?' which means the
first character must be a lowercase letter, and all following
characters must be a dash, lowercase letter, or digit, except the last
character, which cannot be a dash.`,
			},
			"region": {
				Type:             schema.TypeString,
				Required:         true,
				ForceNew:         true,
				DiffSuppressFunc: compareSelfLinkOrResourceName,
				Description:      `The region where the regional SSL policy resides.`,
			},
			"custom_features": {
				Type:     schema.TypeSet,
				Optional: true,
				Description: `A list of features enabled when the selected profile is CUSTOM. The
method returns the set of features that can be specified in this
list. This field must be empty if the profile is not CUSTOM.

See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
for which ciphers are available to use. **Note**: this argument
*must* be present when using the 'CUSTOM' profile. This argument
*must not* be present when using any other profile.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"description": {
				Type:        schema.TypeString,
				Optional:    true,
				ForceNew:    true,
				Description: `An optional description of this resource.`,
			},
			"min_tls_version": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"TLS_1_0", "TLS_1_1", "TLS_1_2", ""}),
				Description: `The minimum version of SSL protocol that can be used by the clients
to establish a connection with the load balancer. Default value: "TLS_1_0" Possible values: ["TLS_1_0", "TLS_1_1", "TLS_1_2"]`,
				Default: "TLS_1_0",
			},
			"profile": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validateEnum([]string{"COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM", ""}),
				Description: `Profile specifies the set of SSL features that can be used by the
load balancer when negotiating SSL with clients. If using 'CUSTOM',
the set of SSL features to enable must be specified in the
'customFeatures' field.

See the [official documentation](https://cloud.google.com/compute/docs/load-balancing/ssl-policies#profilefeaturesupport)
for information on what cipher suites each profile provides. If
'CUSTOM' is used, the 'custom_features' attribute **must be set**. Default value: "COMPATIBLE" Possible values: ["COMPATIBLE", "MODERN", "RESTRICTED", "CUSTOM"]`,
				Default: "COMPATIBLE",
			},
			"creation_timestamp": {
				Type:        schema.TypeString,
				Computed:    true,
				Description: `Creation timestamp in RFC3339 text format.`,
			},
			"enabled_features": {
				Type:        schema.TypeSet,
				Computed:    true,
				Description: `The list of features enabled in the SSL policy.`,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"fingerprint": {
				Type:     schema.TypeString,
				Computed: true,
				Description: `Fingerprint of this resource. A hash of the contents stored in this
object. This field is used in optimistic locking.`,
			},
			"project": {
				Type:     schema.TypeString,
				Optional: true,
				Computed: true,
				ForceNew: true,
			},
			"self_link": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
		UseJSONNumber: true,
	}
}

func resourceComputeRegionSslPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	obj := make(map[string]interface{})
	descriptionProp, err := expandComputeRegionSslPolicyDescription(d.Get("description"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}
	nameProp, err := expandComputeRegionSslPolicyName(d.Get("name"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("name"); !isEmptyValue(reflect.ValueOf(nameProp)) && (ok || !reflect.DeepEqual(v, nameProp)) {
		obj["name"] = nameProp
	}
	profileProp, err := expandComputeRegionSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("profile"); !isEmptyValue(reflect.ValueOf(profileProp)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeRegionSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_tls_version"); !isEmptyValue(reflect.ValueOf(minTlsVersionProp)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeRegionSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_features"); ok || !reflect.DeepEqual(v, customFeaturesProp) {
		obj["customFeatures"] = customFeaturesProp
	}
	fingerprintProp, err := expandComputeRegionSslPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(fingerprintProp)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}
	regionProp, err := expandComputeRegionSslPolicyRegion(d.Get("region"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("region"); !isEmptyValue(reflect.ValueOf(regionProp)) && (ok || !reflect.DeepEqual(v, regionProp)) {
		obj["region"] = regionProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/sslPolicies")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Creating new RegionSslPolicy: %#v", obj)
	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSslPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "POST", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutCreate))
	if err != nil {
		return fmt.Errorf("Error creating RegionSslPolicy: %s", err)
	}

	// Store the ID now
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}")
	if err != nil {
		return fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	err = computeOperationWaitTime(
		config, res, project, "Creating RegionSslPolicy", userAgent,
		d.Timeout(schema.TimeoutCreate))

	if err != nil {
		// The resource didn't actually create
		d.SetId("")
		return fmt.Errorf("Error waiting to create RegionSslPolicy: %s", err)
	}

	log.Printf("[DEBUG] Finished creating RegionSslPolicy %q: %#v", d.Id(), res)

	return resourceComputeRegionSslPolicyRead(d, meta)
}

func resourceComputeRegionSslPolicyRead(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSslPolicy: %s", err)
	}
	billingProject = project

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequest(config, "GET", billingProject, url, userAgent, nil)
	if err != nil {
		return handleNotFoundError(err, d, fmt.Sprintf("ComputeRegionSslPolicy %q", d.Id()))
	}

	if err := d.Set("project", project); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}

	if err := d.Set("creation_timestamp", flattenComputeRegionSslPolicyCreationTimestamp(res["creationTimestamp"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("description", flattenComputeRegionSslPolicyDescription(res["description"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("name", flattenComputeRegionSslPolicyName(res["name"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("profile", flattenComputeRegionSslPolicyProfile(res["profile"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("min_tls_version", flattenComputeRegionSslPolicyMinTlsVersion(res["minTlsVersion"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("enabled_features", flattenComputeRegionSslPolicyEnabledFeatures(res["enabledFeatures"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("custom_features", flattenComputeRegionSslPolicyCustomFeatures(res["customFeatures"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("fingerprint", flattenComputeRegionSslPolicyFingerprint(res["fingerprint"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("region", flattenComputeRegionSslPolicyRegion(res["region"], d, config)); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}
	if err := d.Set("self_link", ConvertSelfLinkToV1(res["selfLink"].(string))); err != nil {
		return fmt.Errorf("Error reading RegionSslPolicy: %s", err)
	}

	return nil
}

func resourceComputeRegionSslPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSslPolicy: %s", err)
	}
	billingProject = project

	obj := make(map[string]interface{})
	profileProp, err := expandComputeRegionSslPolicyProfile(d.Get("profile"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("profile"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, profileProp)) {
		obj["profile"] = profileProp
	}
	minTlsVersionProp, err := expandComputeRegionSslPolicyMinTlsVersion(d.Get("min_tls_version"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("min_tls_version"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, minTlsVersionProp)) {
		obj["minTlsVersion"] = minTlsVersionProp
	}
	customFeaturesProp, err := expandComputeRegionSslPolicyCustomFeatures(d.Get("custom_features"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("custom_features"); ok || !reflect.DeepEqual(v, customFeaturesProp) {
		obj["customFeatures"] = customFeaturesProp
	}
	fingerprintProp, err := expandComputeRegionSslPolicyFingerprint(d.Get("fingerprint"), d, config)
	if err != nil {
		return err
	} else if v, ok := d.GetOkExists("fingerprint"); !isEmptyValue(reflect.ValueOf(v)) && (ok || !reflect.DeepEqual(v, fingerprintProp)) {
		obj["fingerprint"] = fingerprintProp
	}

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Updating RegionSslPolicy %q: %#v", d.Id(), obj)

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "PATCH", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return fmt.Errorf("Error updating RegionSslPolicy %q: %s", d.Id(), err)
	} else {
		log.Printf("[DEBUG] Finished updating RegionSslPolicy %q: %#v", d.Id(), res)
	}

	err = computeOperationWaitTime(
		config, res, project, "Updating RegionSslPolicy", userAgent,
		d.Timeout(schema.TimeoutUpdate))

	if err != nil {
		return err
	}

	return resourceComputeRegionSslPolicyRead(d, meta)
}

func resourceComputeRegionSslPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	config := meta.(*Config)
	userAgent, err := generateUserAgentString(d, config.userAgent)
	if err != nil {
		return err
	}

	billingProject := ""

	project, err := getProject(d, config)
	if err != nil {
		return fmt.Errorf("Error fetching project for RegionSslPolicy: %s", err)
	}
	billingProject = project

	url, err := replaceVars(d, config, "{{ComputeBasePath}}projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}")
	if err != nil {
		return err
	}

	var obj map[string]interface{}
	log.Printf("[DEBUG] Deleting RegionSslPolicy %q", d.Id())

	// err == nil indicates that the billing_project value was found
	if bp, err := getBillingProject(d, config); err == nil {
		billingProject = bp
	}

	res, err := sendRequestWithTimeout(config, "DELETE", billingProject, url, userAgent, obj, d.Timeout(schema.TimeoutDelete))
	if err != nil {
		return handleNotFoundError(err, d, "RegionSslPolicy")
	}

	err = computeOperationWaitTime(
		config, res, project, "Deleting RegionSslPolicy", userAgent,
		d.Timeout(schema.TimeoutDelete))

	if err != nil {
		return err
	}

	log.Printf("[DEBUG] Finished deleting RegionSslPolicy %q: %#v", d.Id(), res)
	return nil
}

func resourceComputeRegionSslPolicyImport(d *schema.ResourceData, meta interface{}) ([]*schema.ResourceData, error) {
	config := meta.(*Config)
	if err := parseImportId([]string{
		"projects/(?P<project>[^/]+)/regions/(?P<region>[^/]+)/sslPolicies/(?P<name>[^/]+)",
		"(?P<project>[^/]+)/(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<region>[^/]+)/(?P<name>[^/]+)",
		"(?P<name>[^/]+)",
	}, d, config); err != nil {
		return nil, err
	}

	// Replace import id for the resource id
	id, err := replaceVars(d, config, "projects/{{project}}/regions/{{region}}/sslPolicies/{{name}}")
	if err != nil {
		return nil, fmt.Errorf("Error constructing id: %s", err)
	}
	d.SetId(id)

	return []*schema.ResourceData{d}, nil
}

func flattenComputeRegionSslPolicyCreationTimestamp(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionSslPolicyDescription(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionSslPolicyName(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionSslPolicyProfile(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionSslPolicyMinTlsVersion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionSslPolicyEnabledFeatures(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeRegionSslPolicyCustomFeatures(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return schema.NewSet(schema.HashString, v.([]interface{}))
}

func flattenComputeRegionSslPolicyFingerprint(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	return v
}

func flattenComputeRegionSslPolicyRegion(v interface{}, d *schema.ResourceData, config *Config) interface{} {
	if v == nil {
		return v
	}
	return ConvertSelfLinkToV1(v.(string))
}

func expandComputeRegionSslPolicyDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyProfile(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyMinTlsVersion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyCustomFeatures(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	v = v.(*schema.Set).List()
	return v, nil
}

func expandComputeRegionSslPolicyFingerprint(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandComputeRegionSslPolicyRegion(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	f, err := parseGlobalFieldValue("regions", v.(string), "project", d, config, true)
	if err != nil {
		return nil, fmt.Errorf("Invalid value for region: %s", err)
	}
	return f.RelativeLink(), nil
}
