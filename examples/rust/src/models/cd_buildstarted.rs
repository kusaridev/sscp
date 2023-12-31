/*
 * Secure Build API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 * Generated by: https://openapi-generator.tech
 */




#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct CdBuildstarted {
    #[serde(rename = "context")]
    pub context: Box<crate::models::CdBuildstartedContext>,
    #[serde(rename = "subject")]
    pub subject: Box<crate::models::CdBuildstartedSubject>,
    #[serde(rename = "customData", skip_serializing_if = "Option::is_none")]
    pub custom_data: Option<Box<crate::models::CdBuildstartedCustomData>>,
    #[serde(rename = "customDataContentType", skip_serializing_if = "Option::is_none")]
    pub custom_data_content_type: Option<String>,
}

impl CdBuildstarted {
    pub fn new(context: crate::models::CdBuildstartedContext, subject: crate::models::CdBuildstartedSubject) -> CdBuildstarted {
        CdBuildstarted {
            context: Box::new(context),
            subject: Box::new(subject),
            custom_data: None,
            custom_data_content_type: None,
        }
    }
}


