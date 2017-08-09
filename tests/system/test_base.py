from axtx_beats import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Axtx_beats normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        axtx_beats_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("axtx_beats is running"))
        exit_code = axtx_beats_proc.kill_and_wait()
        assert exit_code == 0
