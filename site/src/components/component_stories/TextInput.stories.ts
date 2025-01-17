import type { Meta, StoryObj } from '@storybook/vue3'

import TextInput from '@/components/TextInput.vue'

const meta: Meta<typeof TextInput> = {
  component: TextInput,
}

export default meta
type Story = StoryObj<typeof TextInput>

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { TextInput },
    setup() {
      return { args }
    },
    template: '<TextInput v-bind="args"/>',
  }),
  args: {
    placeholder: "Username",
  },
}
