import * as React from "react";
import { useForm } from "react-hook-form";

import useChess from "../stores/chess";

interface ChessStartProps {
  onAuthorize: () => void;
}

const ChessStart: React.FC<ChessStartProps> = ({ onAuthorize }) => {
  const { play } = useChess((s) => ({ play: s.play }));
  const {
    handleSubmit,
    formState: { errors },
    setError,
    register,
  } = useForm<{ hours: number; minutes: number }>({
    reValidateMode: "onSubmit",
    mode: "onSubmit",
  });

  return (
    <div className="flex flex-1 items-center justify-center flex-col">
      <div className="row flex items-center my-10">
        <label className="text-sm font-medium text-gray-900 dark:text-white">
          Per iniziare la partita seleziona la durata di ogni turno:
        </label>
      </div>
      <form
        className="flex flex-col"
        onSubmit={handleSubmit((data) => {
          if (data.minutes == 0 && data.hours == 0) {
            setError("minutes", {
              message: "Il turno deve durare almeno 1 minuto",
            });
          } else {
            play(data);
            onAuthorize();
          }
        })}
      >
        <div className="w-full flex flex-row justify-center items-center">
          <div className="flex flex-col items-start mx-4">
            <label className="block mb-2 text-sm text-center font-medium text-gray-900 dark:text-white">
              Ore
            </label>
            <input
              type="number"
              placeholder="Ore"
              className="w-full mb-2 rounded-lg text-center dark:bg-gray-700"
              max={100}
              min={0}
              defaultValue={0}
              {...register("hours", {
                required: true,
                valueAsNumber: true,
              })}
            />
          </div>

          <div className="flex flex-col items-start mx-4">
            <label className="block mb-2 text-sm text-center font-medium text-gray-900 dark:text-white">
              Minuti
            </label>
            <input
              type="number"
              placeholder="Minuti"
              className="w-full mb-2 rounded-lg text-center dark:bg-gray-700"
              max={59}
              min={0}
              defaultValue={1}
              {...register("minutes", {
                required: true,
                valueAsNumber: true,
              })}
            />
          </div>
        </div>
        {errors.minutes && (
          <label htmlFor="hours">
            {errors.minutes.message ||
              "I minuti devono essere compresi tra 0 e 59"}
          </label>
        )}
        {errors.hours && (
          <label htmlFor="hours">
            Le ore devono essere comprese tra 0 e 100
          </label>
        )}
        <div className="flex justify-center">
          <button
            type="submit"
            className="w-full mt-5 justify-center text-white text-center hover:bg-sky-700 bg-sky-700 hover:bg-sky-800 focus:ring-4 focus:outline-none focus:ring-sky-300 font-medium rounded-lg text-sm px-4 py-2 dark:bg-sky-600 dark:hover:bg-sky-700 dark:focus:ring-sky-800"
          >
            Inizia la partita!
          </button>
        </div>
      </form>
    </div>
  );
};

export default ChessStart;
