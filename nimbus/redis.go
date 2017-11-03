package nimbus

import (
	. "github.com/cloudfoundry/cf-acceptance-tests/cats_suite_helpers"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gexec"

	"github.com/cloudfoundry-incubator/cf-test-helpers/cf"
	"github.com/cloudfoundry-incubator/cf-test-helpers/helpers"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers/app_helpers"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers/assets"
	"github.com/cloudfoundry/cf-acceptance-tests/helpers/random_name"
)

var _ = NimbusDescribe("redis service", func() {

	var appName, redisName string

	BeforeEach(func() {

		if Config.GetIncludeNimbusServiceRedis() != true {
			Skip("include_nimbus_service_redis was not set to true")
		}

		appName = random_name.CATSRandomName("APP")
		redisName = random_name.CATSRandomName("SVC")

		Expect(cf.Cf("create-service", Config.GetNimbusServiceNameRedis(), "shared-vm", redisName).Wait(Config.DefaultTimeoutDuration())).To(Exit(0))

		Expect(cf.Cf("push", appName, "-p", assets.NewAssets().NimbusServices, "--no-start", "-i", "2").Wait(Config.CfPushTimeoutDuration())).To(Exit(0))
		Expect(cf.Cf("bind-service", appName, redisName).Wait(Config.DefaultTimeoutDuration())).To(Exit(0))

		app_helpers.SetBackend(appName)
		Expect(cf.Cf("start", appName).Wait(Config.CfPushTimeoutDuration())).To(Exit(0))
	})

	AfterEach(func() {
		Expect(cf.Cf("delete", appName, "-f").Wait(Config.DefaultTimeoutDuration())).To(Exit(0))
		Expect(cf.Cf("delete-service", redisName, "-f").Wait(Config.DefaultTimeoutDuration())).To(Exit(0))
	})

	It("is accessible in datacenters", func() {

		randomKey := random_name.CATSRandomName("KEY")
		randomValue := random_name.CATSRandomName("VAL")

		Eventually(func() string {
			return helpers.CurlApp(Config, appName, "/redis/insert/"+randomKey+"/"+randomValue)
		}, Config.DefaultTimeoutDuration()).Should(ContainSubstring("OK"))

		Eventually(func() string {
			return helpers.CurlApp(Config, appName, "/redis/read/"+randomKey+"/"+randomValue)
		}, Config.DefaultTimeoutDuration()).Should(ContainSubstring("OK"))

	})

})