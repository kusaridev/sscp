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
pub struct CdBuildstartedContext {
    #[serde(rename = "version")]
    pub version: String,
    #[serde(rename = "id")]
    pub id: String,
    #[serde(rename = "source")]
    pub source: String,
    #[serde(rename = "type")]
    pub r#type: Type,
    #[serde(rename = "timestamp")]
    pub timestamp: String,
}

impl CdBuildstartedContext {
    pub fn new(version: String, id: String, source: String, r#type: Type, timestamp: String) -> CdBuildstartedContext {
        CdBuildstartedContext {
            version,
            id,
            source,
            r#type,
            timestamp,
        }
    }
}

/// 
#[derive(Clone, Copy, Debug, Eq, PartialEq, Ord, PartialOrd, Hash, Serialize, Deserialize)]
pub enum Type {
    #[serde(rename = "dev.cdevents.build.started.0.1.1")]
    DevPeriodCdeventsPeriodBuildPeriodStartedPeriod0Period1Period1,
}

impl Default for Type {
    fn default() -> Type {
        Self::DevPeriodCdeventsPeriodBuildPeriodStartedPeriod0Period1Period1
    }
}
