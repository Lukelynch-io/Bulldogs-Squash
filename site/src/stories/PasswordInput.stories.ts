import type { Meta, StoryObj } from '@storybook/vue3'

import PasswordInput from '@/components/PasswordInput.vue'

const meta: Meta<typeof PasswordInput> = {
  component: PasswordInput,
}

export default meta
type Story = StoryObj<typeof PasswordInput>

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { PasswordInput },
    setup() {
      return { args }
    },
    template: '<PasswordInput v-bind="args"/>',
  }),
  args: {
    placeholder: "Password",
  },
}
