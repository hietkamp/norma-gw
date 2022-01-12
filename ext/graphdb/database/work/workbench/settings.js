{
  "users" : {
    "admin" : {
      "username" : "admin",
      "password" : "{bcrypt}$2a$10$bk5AhkSNCvm/EyBIdZA9m.neE92S5UII1d0ZeHWX0cwTjp5hbxF2K",
      "grantedAuthorities" : [ "ROLE_ADMIN" ],
      "appSettings" : {
        "DEFAULT_INFERENCE" : true,
        "DEFAULT_VIS_GRAPH_SCHEMA" : true,
        "DEFAULT_SAMEAS" : true,
        "IGNORE_SHARED_QUERIES" : false,
        "EXECUTE_COUNT" : true
      },
      "dateCreated" : 1641486791036
    }
  },
  "import.server" : { },
  "import.local" : {
    "news;;pub-properties.ttl" : {
      "name" : "pub-properties.ttl",
      "status" : "DONE",
      "message" : "Imported successfully in less than a second.",
      "context" : "",
      "replaceGraphs" : [ ],
      "baseURI" : "file:/uploaded/generated/pub-properties.ttl",
      "forceSerial" : false,
      "type" : "file",
      "format" : null,
      "data" : "504417ad-7009-4e25-9a28-5ce666dbe330",
      "timestamp" : 1641798235330,
      "parserSettings" : {
        "preserveBNodeIds" : false,
        "failOnUnknownDataTypes" : false,
        "verifyDataTypeValues" : false,
        "normalizeDataTypeValues" : false,
        "failOnUnknownLanguageTags" : false,
        "verifyLanguageTags" : true,
        "normalizeLanguageTags" : false,
        "stopOnError" : true
      },
      "requestIdHeadersToForward" : {
        "X-Request-Id" : "c2873e84-bb8a-5204-ab64-b260df4a9ce6"
      }
    },
    "news;;pub-ontology.ttl" : {
      "name" : "pub-ontology.ttl",
      "status" : "DONE",
      "message" : "Imported successfully in less than a second.",
      "context" : "",
      "replaceGraphs" : [ ],
      "baseURI" : "file:/uploaded/generated/pub-ontology.ttl",
      "forceSerial" : false,
      "type" : "file",
      "format" : null,
      "data" : "28c4b993-6d19-49a4-aa61-b27b1fa090c9",
      "timestamp" : 1641798235345,
      "parserSettings" : {
        "preserveBNodeIds" : false,
        "failOnUnknownDataTypes" : false,
        "verifyDataTypeValues" : false,
        "normalizeDataTypeValues" : false,
        "failOnUnknownLanguageTags" : false,
        "verifyLanguageTags" : true,
        "normalizeLanguageTags" : false,
        "stopOnError" : true
      },
      "requestIdHeadersToForward" : {
        "X-Request-Id" : "e73c5dfb-8658-5896-af42-ba16209cff04"
      }
    },
    "news;;pub-ontology-types.ttl" : {
      "name" : "pub-ontology-types.ttl",
      "status" : "DONE",
      "message" : "Imported successfully in less than a second.",
      "context" : "",
      "replaceGraphs" : [ ],
      "baseURI" : "file:/uploaded/generated/pub-ontology-types.ttl",
      "forceSerial" : false,
      "type" : "file",
      "format" : null,
      "data" : "6c3ebbd5-5d69-4d41-9cec-e6d38a17de74",
      "timestamp" : 1641798235371,
      "parserSettings" : {
        "preserveBNodeIds" : false,
        "failOnUnknownDataTypes" : false,
        "verifyDataTypeValues" : false,
        "normalizeDataTypeValues" : false,
        "failOnUnknownLanguageTags" : false,
        "verifyLanguageTags" : true,
        "normalizeLanguageTags" : false,
        "stopOnError" : true
      },
      "requestIdHeadersToForward" : {
        "X-Request-Id" : "8fb78b8b-1a8c-52ee-adea-a9e9e51bf454"
      }
    },
    "news;;publishing-ontology.ttl" : {
      "name" : "publishing-ontology.ttl",
      "status" : "DONE",
      "message" : "Imported successfully in 1s.",
      "context" : "",
      "replaceGraphs" : [ ],
      "baseURI" : "file:/uploaded/generated/publishing-ontology.ttl",
      "forceSerial" : false,
      "type" : "file",
      "format" : null,
      "data" : "14d1a05e-c2fb-40ce-a65f-f2fd0a8f1451",
      "timestamp" : 1641798235384,
      "parserSettings" : {
        "preserveBNodeIds" : false,
        "failOnUnknownDataTypes" : false,
        "verifyDataTypeValues" : false,
        "normalizeDataTypeValues" : false,
        "failOnUnknownLanguageTags" : false,
        "verifyLanguageTags" : true,
        "normalizeLanguageTags" : false,
        "stopOnError" : true
      },
      "requestIdHeadersToForward" : {
        "X-Request-Id" : "6c6c4bdd-5f33-550b-a19d-abd2b45c4d2d"
      }
    },
    "news;;graphdb-news-dataset.nt" : {
      "name" : "graphdb-news-dataset.nt",
      "status" : "DONE",
      "message" : "Imported successfully in 1m 52s.",
      "context" : "",
      "replaceGraphs" : [ ],
      "baseURI" : "file:/uploaded/generated/graphdb-news-dataset.nt",
      "forceSerial" : false,
      "type" : "file",
      "format" : null,
      "data" : "8f98fdd2-7f1f-4cbf-85b6-291ddab4ea19",
      "timestamp" : 1641798236327,
      "parserSettings" : {
        "preserveBNodeIds" : false,
        "failOnUnknownDataTypes" : false,
        "verifyDataTypeValues" : false,
        "normalizeDataTypeValues" : false,
        "failOnUnknownLanguageTags" : false,
        "verifyLanguageTags" : true,
        "normalizeLanguageTags" : false,
        "stopOnError" : true
      },
      "requestIdHeadersToForward" : {
        "X-Request-Id" : "9ff3c382-eb35-555f-8417-9b0279b98cc2"
      }
    }
  },
  "properties" : {
    "current.location" : ""
  },
  "user_queries" : {
    "admin" : {
      "SPARQL Select template" : {
        "name" : "SPARQL Select template",
        "body" : "SELECT ?s ?p ?o\nWHERE {\n\t?s ?p ?o .\n} LIMIT 100",
        "shared" : false
      },
      "Clear graph" : {
        "name" : "Clear graph",
        "body" : "CLEAR GRAPH <http://example>",
        "shared" : false
      },
      "Add statements" : {
        "name" : "Add statements",
        "body" : "PREFIX dc: <http://purl.org/dc/elements/1.1/>\nINSERT DATA\n      {\n      GRAPH <http://example> {\n          <http://example/book1> dc:title \"A new book\" ;\n                                 dc:creator \"A.N.Other\" .\n          }\n      }",
        "shared" : false
      },
      "Remove statements" : {
        "name" : "Remove statements",
        "body" : "PREFIX dc: <http://purl.org/dc/elements/1.1/>\nDELETE DATA\n{\nGRAPH <http://example> {\n    <http://example/book1> dc:title \"A new book\" ;\n                           dc:creator \"A.N.Other\" .\n    }\n}",
        "shared" : false
      }
    }
  }
}