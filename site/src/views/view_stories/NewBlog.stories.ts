import type { Meta, StoryObj } from '@storybook/vue3';
import NewBlog from '../NewBlog.vue';

const meta: Meta<typeof NewBlog> = {
  component: NewBlog,
};

export default meta;
type Story = StoryObj<typeof NewBlog>;

export const Primary: Story = {
  render: (args) => ({
    components: { NewBlog: NewBlog },
    setup() {
      return { args }
    },
    template: '<NewBlog v-bind="args"/>',
  }),
  args: {
  },
}
