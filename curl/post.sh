#!/usr/bin/env bash

# Example post command.
curl \
-H "Content-Type: application/json" \
-i \
-X POST \
-d '
  {
    "Id": "464a4165-842a-4279-8db6-e47232d2bc60",
    "MajorLabel": "0",
    "MajorVersion": 0,
    "Metadata": "",
    "MinorVersion": 0,
    "RevisionNotes": "",
    "RevisionReason": "",
    "S3URL": "s3://dp-csv-splitter/swooosh/big-all.csv",
    "Status": "complete",
    "Title": "big-all.csv",
    "TotalRowCount": 33732,
    "DataResource": ""
  }' \
http://localhost:8000/dataset
