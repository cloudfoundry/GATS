package quotas_test

import(
	"github.com/nu7hatch/gouuid"
	
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gbytes"
	. "github.com/pivotal-cf-experimental/cf-test-helpers/cf"
)

var assertionTimeout = 10.0

var _ = Describe("CF Quota commands", func() {
	It("can Create, Read, Update, and Delete quotas", func() {
		AsUser(context.AdminUserContext(), func() {
			quotaBytes, err := uuid.NewV4()
			Expect(err).ToNot(HaveOccurred())
			quotaName := quotaBytes.String()
			
			Eventually(Cf("create-quota",
				quotaName,
				"-m", "512M",
			), assertionTimeout).Should(Say("OK"))

			Eventually(Cf("quota", quotaName), assertionTimeout).Should(Say("512M"))

			quotaOutput := Cf("quotas")
			Eventually(quotaOutput, assertionTimeout).Should(Say(quotaName))

			Eventually(Cf("update-quota",
				quotaName,
				"-m", "513M",
			), assertionTimeout).Should(Say("OK"))

			Eventually(Cf("quotas"), assertionTimeout).Should(Say("513M"))

			Eventually(Cf("delete-quota",
				quotaName,
				"-f",
			), assertionTimeout).Should(Say("OK"))

			Eventually(Cf("quotas"), assertionTimeout).ShouldNot(Say(quotaName))
		})
	})
})
