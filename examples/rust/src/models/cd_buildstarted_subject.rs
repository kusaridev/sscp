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
pub struct CdBuildstartedSubject {
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "source", skip_serializing_if = "Option::is_none")]
    pub source: Option<String>,
    #[serde(rename = "type")]
    pub r#type: Type,
    #[serde(rename = "content")]
    pub content: serde_json::Value,
}

impl CdBuildstartedSubject {
    pub fn new(id: String, r#type: Type, content: serde_json::Value) -> CdBuildstartedSubject {
        CdBuildstartedSubject {
            id,
            source: None,
            r#type,
            content,
        }
    }
}

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Type {
    #[serde(rename = "build")]
    Build,
}

impl Default for Type {
    fn default() -> Type {
        Self::Build
    }
}

