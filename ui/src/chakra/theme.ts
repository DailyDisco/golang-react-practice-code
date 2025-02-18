import { createSystem, defaultConfig } from '@chakra-ui/react';

export const system = createSystem(defaultConfig, {
  theme: {
    tokens: {
      colors: {
        primary: { value: '#3182CE' },
        secondary: { value: '#718096' },
        success: { value: '#38A169' },
        error: { value: '#E53E3E' },
        background: {
          light: { value: '#F7FAFC' },
          dark: { value: '#1A202C' },
        },
        surface: {
          light: { value: '#FFFFFF' },
          dark: { value: '#2D3748' },
        },
        text: {
          light: { value: '#2D3748' },
          dark: { value: '#F7FAFC' },
        },
        muted: {
          light: { value: '#718096' },
          dark: { value: '#A0AEC0' },
        },
        border: {
          light: { value: '#E2E8F0' },
          dark: { value: '#4A5568' },
        },
        hover: {
          light: { value: '#EDF2F7' },
          dark: { value: '#2D3748' },
        },
      },
      fonts: {
        heading: { value: `'Figtree', sans-serif` },
        body: { value: `'Figtree', sans-serif` },
      },
      fontSizes: {
        xs: { value: '0.75rem' },
        sm: { value: '0.875rem' },
        md: { value: '1rem' },
        lg: { value: '1.125rem' },
        xl: { value: '1.25rem' },
        '2xl': { value: '1.5rem' },
      },
      radii: {
        sm: { value: '0.375rem' },
        md: { value: '0.5rem' },
        lg: { value: '0.75rem' },
        full: { value: '9999px' },
      },
      shadows: {
        sm: { value: '0 1px 3px 0 rgba(0, 0, 0, 0.1)' },
        md: { value: '0 4px 6px -1px rgba(0, 0, 0, 0.1)' },
        lg: { value: '0 10px 15px -3px rgba(0, 0, 0, 0.1)' },
      },
      spacing: {
        xs: { value: '0.5rem' },
        sm: { value: '1rem' },
        md: { value: '1.5rem' },
        lg: { value: '2rem' },
        xl: { value: '3rem' },
      },
    },
    semanticTokens: {},
  },
});

export default system;
