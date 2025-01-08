import type { Meta, StoryObj } from '@storybook/vue3'

import TextArea from '@/components/TextArea.vue'

const meta: Meta<typeof TextArea> = {
  component: TextArea,
}

export default meta
type Story = StoryObj<typeof TextArea>

/*
 *👇 Render functions are a framework specific feature to allow you control on how the component renders.
 * See https://storybook.js.org/docs/api/csf
 * to learn how to use render functions.
 */
export const Primary: Story = {
  render: (args) => ({
    components: { TextArea: TextArea },
    setup() {
      return { args }
    },
    template: '<TextArea v-bind="args"/>',
  }),
  args: {
    placeholder: "Post Detail",
  },
}
