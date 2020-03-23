import groovy.transform.Field

@Field String STEP_NAME = getClass().getName()
@Field String METADATA_FILE = 'metadata/nexusUpload.yaml'

//Metadata maintained in file project://resources/metadata/nexusUpload.yaml

void call(Map parameters = [:]) {
    List credentials = [[type: 'usernamePassword', id: 'nexusCredentialsId', env: ['PIPER_username', 'PIPER_password']]]
    piperExecuteBin(parameters, STEP_NAME, METADATA_FILE, credentials)
}
