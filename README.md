ECSLOGS
==============

[![wercker status](https://app.wercker.com/status/6ee0a839467f6209828fb2e748383b71/s/master "wercker status")](https://app.wercker.com/project/byKey/6ee0a839467f6209828fb2e748383b71)

OVERVIEW
--------------

ECSLOGS is a tool developped in Golang to gather logs files from an ECS VDC.

BUILD
--------------

The Dockerfile can be used to create a Docker container for this tool.

Just run the following command in the folder that contains the Dockerfile: docker build -t ecslogs .

RUN
--------------

```
echo <password> | docker run -i djannot/ecslogs ./ecslogs <user> <host:port> <pattern> <input log> <# days> <pipe|file> <output file>
```

For example, to get the access logs for the last 3 days:
```
 echo "ChangeMe" | docker run -i -v `pwd`:/tmp djannot/ecslogs ./ecslogs admin 10.64.231.196:22 RequestLog.java dataheadsvc.log 3 file /tmp/access.log
```

It will automatically get the logs from all the ECS nodes of the VDC.

LICENSING
--------------

Licensed under the Apache License, Version 2.0 (the “License”); you may not use this file except in compliance with the License. You may obtain a copy of the License at <http://www.apache.org/licenses/LICENSE-2.0>

Unless required by applicable law or agreed to in writing, software distributed under the License is distributed on an “AS IS” BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied. See the License for the specific language governing permissions and limitations under the License.
