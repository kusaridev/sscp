/*
 * Secure Build API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * The version of the OpenAPI document: 0.0.1
 * 
 * Generated by: https://openapi-generator.tech
 */

/// Subject : Represents a subject in an In-Toto v1 statement.



#[derive(Clone, Debug, PartialEq, Serialize, Deserialize)]
pub struct Subject {
    /// Represents a set of digests, mapping algorithms to their respective digest strings.
    #[serde(rename = "digest")]
    pub digest: ::std::collections::HashMap<String, String>,
    #[serde(rename = "name")]
    pub name: String,
}

impl Subject {
    /// Represents a subject in an In-Toto v1 statement.
    pub fn new(digest: ::std::collections::HashMap<String, String>, name: String) -> Subject {
        Subject {
            digest,
            name,
        }
    }
}


