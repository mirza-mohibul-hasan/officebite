import { AlertCircle } from 'lucide-react';

type ErrorStateProps = {
  title?: string;
  message: string;
};

export function ErrorState({ title = 'Something went wrong', message }: ErrorStateProps) {
  return (
    <div className="flex gap-3 rounded-lg border border-red-200 bg-red-50 p-4 text-sm text-red-900">
      <AlertCircle className="mt-0.5 shrink-0" size={18} aria-hidden="true" />
      <div>
        <h3 className="font-semibold">{title}</h3>
        <p className="mt-1 leading-6">{message}</p>
      </div>
    </div>
  );
}
