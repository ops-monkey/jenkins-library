import com.sap.piper.ConfigurationHelper
import com.sap.piper.GenerateStageDocumentation
import com.sap.piper.ReportAggregator
import com.sap.piper.Utils
import groovy.transform.Field

import static com.sap.piper.Prerequisites.checkScript

@Field String STEP_NAME = getClass().getName()

@Field Set GENERAL_CONFIG_KEYS = []
@Field Set STAGE_STEP_KEYS = [
    /** Parameters for deployment to a Nexus Repository Manager. */
    'nexus'
]
@Field Set STEP_CONFIG_KEYS = GENERAL_CONFIG_KEYS.plus(STAGE_STEP_KEYS)
@Field Set PARAMETER_KEYS = STEP_CONFIG_KEYS

/**
 * This stage is responsible for releasing/deploying artifacts to a Nexus Repository Manager.<br />
 */
@GenerateStageDocumentation(defaultStageName = 'artifactDeployment')
void call(Map parameters = [:]) {
    String stageName = 'artifactDeployment'
    final script = checkScript(this, parameters) ?: this
    def utils = parameters.juStabUtils ?: new Utils()

    Map config = ConfigurationHelper.newInstance(this)
        .loadStepDefaults()
        .mixinGeneralConfig(script.commonPipelineEnvironment, GENERAL_CONFIG_KEYS)
        .mixinStageConfig(script.commonPipelineEnvironment, stageName, STEP_CONFIG_KEYS)
        .mixin(parameters, PARAMETER_KEYS)
        .withMandatoryProperty('nexus')
        .use()

    piperStageWrapper(stageName: stageName, script: script) {

        def commonPipelineEnvironment = script.commonPipelineEnvironment
        List unstableSteps = commonPipelineEnvironment?.getValue('unstableSteps') ?: []
        if (unstableSteps) {
            piperPipelineStageConfirm script: script
            unstableSteps = []
            commonPipelineEnvironment.setValue('unstableSteps', unstableSteps)
        }

        // telemetry reporting
        utils.pushToSWA([step: STEP_NAME], config)

        Map nexusConfig = config.nexus as Map

        // Add all mandatory parameters
        Map nexusUploadParams = [
            script: script,
//            version: nexusConfig.version,
//            repository: nexusConfig.repository,
//            url: nexusConfig.url,
            additionalClassifiers: nexusConfig.additionalClassifiers,
        ]

        // Set artifactId from CPE if set
        if (script.commonPipelineEnvironment.configuration.artifactId) {
            nexusUploadParams.artifactId = script.commonPipelineEnvironment.configuration.artifactId
        }

        // REPOSITORY_UNDER_TEST and LIBRARY_VERSION_UNDER_TEST have to be removed from withEnv before merging to master.
        withEnv(["STAGE_NAME=${stageName}", 'REPOSITORY_UNDER_TEST=SAP/jenkins-library','LIBRARY_VERSION_UNDER_TEST=stage-artifact-deployment']) {
            nexusUpload(nexusUploadParams)
        }

        ReportAggregator.instance.reportDeploymentToNexus()
    }
}
