import type { Meta, StoryObj } from '@storybook/vue3';
import BlogPostCollection from '../BlogPostCollection.vue';

const meta: Meta<typeof BlogPostCollection> = {
  component: BlogPostCollection,
};

export default meta;
type Story = StoryObj<typeof BlogPostCollection>;

export const Primary: Story = {
  render: (args) => ({
    components: { BlogPostCollection: BlogPostCollection },
    setup() {
      return { args }
    },
    template: '<BlogPostCollection v-bind="args"/>',
  }),
  args: {
  },
}
