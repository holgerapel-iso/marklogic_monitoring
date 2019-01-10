# MarkLogic Monitoring Utility
This utility uses the MarkLogic Management API to obtain the health information.

The received health-reports are filtered by the provided resourceType and states.

The result is written to stdout and can be used to trigger specific monitoring actions.

## Examples 
* Check all resourceTypes with any state other than "info"

  `marklogic_monitor "http://localhost:8002/manage/v2?view=health&format=json" username password`

* Check all resourceTypes with any state other than "info" or "at-risk"

  `marklogic_monitor "http://localhost:8002/manage/v2?view=health&format=json" username password "*" "info,at-risk"` 

* Only check resourceType "database" with any state other than "info"
  
  `marklogic_monitor "http://localhost:8002/manage/v2?view=health&format=json" username password database info`

## Output 
* When no critical healt-report was found
  
  `OK` 
 
* Each matching health-report is shown with their state, resourceType, id, resourceName, Code and Message
  
  `{at-risk cluster 7222958793381141742 localhost-cluster HEALTH-CLUSTER  Review the following critical resources: forest:Triggers_R02}` 
 
