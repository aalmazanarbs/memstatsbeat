from memstatsbeat import BaseTest

import os


class Test(BaseTest):

    def test_base(self):
        """
        Basic test with exiting Memstatsbeat normally
        """
        self.render_config_template(
            path=os.path.abspath(self.working_dir) + "/log/*"
        )

        memstatsbeat_proc = self.start_beat()
        self.wait_until(lambda: self.log_contains("memstatsbeat is running"))
        exit_code = memstatsbeat_proc.kill_and_wait()
        assert exit_code == 0
