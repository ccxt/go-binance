## Generate SBE code
- Install Java
- Download latest sbe schema
```bash
git clone https://github.com/aeron-io/simple-binary-encoding.git
cd simple-binary-encoding
./gradlew
```
Run following command to generate golang code:
```bash
java --add-opens java.base/jdk.internal.misc=ALL-UNNAMED -Dsbe.generate.ir=true -Dsbe.target.language=Go -Dsbe.target.namespace=sbe -Dsbe.output.dir=include/gen -Dsbe.errorLog=yes -jar sbe-all/build/libs/sbe-all-${SBE_TOOL_VERSION}.jar {SBE_SCHEMA}.xml
```


